package repository

import (
	"errors"
	"fmt"
	"opslaundry/pkg/commons"
	"opslaundry/pkg/constants"
	"opslaundry/pkg/utils"
	"reflect"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BaseRepository interface {
	GetPagination(context *gin.Context, request commons.DataTableRequest, isTenantFilter bool, sqlFilter ...interface{}) (interface{}, error)
}

type baseConnection struct {
	DB *gorm.DB
}

func NewBaseRepository(db *gorm.DB) BaseRepository {
	return &baseConnection{
		DB: db,
	}
}

var (
	dateNow time.Time = time.Now().UTC()
)

func getFields(context *gin.Context) []string {
	var output []string

	t := reflect.TypeOf(context.Value("table"))
	for i := 0; i < t.NumField(); i++ {
		switch jsonTag := t.Field(i).Tag.Get("json"); jsonTag {
		case "-":
		case "token,omitempty":
		case "":
			fmt.Println("")
		default:
			//fmt.Println(jsonTag)
			output = append(output, jsonTag)
		}
	}

	v := reflect.TypeOf(context.Value("view"))
	for i := 0; i < v.NumField(); i++ {
		//fmt.Println(v.Field(i))

		switch jsonTag := v.Field(i).Tag.Get("json"); jsonTag {
		case "-":
		case "token,omitempty":
		case "":
			fmt.Println("")
		default:
			//fmt.Println(jsonTag)
			output = append(output, jsonTag)
		}
	}
	return output
}

func fieldInSlice(name string, list []string) bool {
	for _, item := range list {
		if item == name {
			return true
		}
	}
	return false
}

func filterByOperator(field string, operator string, value string) string {
	var result = ""
	var _operator = ""
	var _field = field
	var _value = value
	switch os := operator; os {
	case "isnull":
		_operator = "IS NULL"
	case "notnull":
		_operator = "IS NOT NULL"
	case "neq":
		_operator = "<>"
	case "lt":
		_operator = "<"
	case "lte":
		_operator = "<="
	case "gt":
		_operator = ">"
	case "gte":
		_operator = ">="
	default:
		_operator = "="
		_field = fmt.Sprintf("LOWER(%v::text)", field)
		_value = fmt.Sprintf("LOWER('%v')", value)
	}

	if operator == "isnull" || operator == "isnotnull" {
		result = fmt.Sprintf("%v %v", _field, _operator)
	} else {
		result = fmt.Sprintf("%v %v %v", _field, _operator, _value)
	}
	return result
}

func (r baseConnection) GetPagination(context *gin.Context, request commons.DataTableRequest, isTenantFilter bool, sqlFilter ...interface{}) (interface{}, error) {
	var response commons.DataTableResponse
	tableName, ok := context.Get("table_name")
	if !ok {
		return nil, errors.New("Table name is undefined")
	}

	userLogin, ok := context.Get(constants.USER_IDENTITY)
	if !ok {
		return nil, errors.New("Failed to get UserIdentity from context")
	}

	var records []map[string]interface{}
	var selectedFields []string = getFields(context)
	var conditions string = "1 = 1"

	//region Validation
	if len(request.Filters) > 0 {
		for _, v := range request.Filters {
			if !fieldInSlice(v.Field, selectedFields) {
				return nil, errors.New(fmt.Sprintf("Kolom %v dalam Filter tidak terdaftar.", v.Field))
			}
		}
	}
	if len(request.Orders) > 0 {
		for _, v := range request.Orders {
			if !fieldInSlice(v.Field, selectedFields) {
				return nil, errors.New(fmt.Sprintf("Kolom %v dalam Order tidak terdaftar.", v.Field))
			}
		}
	}
	//endregion

	//-- Filter whole records by IsSystemAdmin
	if userLogin.(utils.UserIdentity).IsSystemAdmin == false {
		conditions = fmt.Sprintf("organization_id::TEXT = '%v'", userLogin.(utils.UserIdentity).OrganizationId)
	}

	if err := r.DB.Raw(fmt.Sprintf("SELECT COUNT(id) FROM %v WHERE %v", tableName, conditions)).Scan(&response.RecordsTotal).Error; err != nil {
		return nil, err
	}

	if response.RecordsTotal > 0 {
		page := request.Page
		if page == 0 {
			page = 1
		}

		pageSize := request.Size
		switch {
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (page - 1) * pageSize

		if request.Search != nil && *request.Search != "" {
			conditions += " AND ("
			total := 0

			for _, v := range selectedFields {
				if total > 0 {
					conditions += " OR "
				}
				conditions += fmt.Sprintf("%v::TEXT ILIKE '%v'", v, "%"+*request.Search+"%")
				total++
			}
			conditions += ")"
		}

		if len(request.Filters) > 0 {
			conditions += " AND ("
			for _, v := range request.Filters {
				total := 0
				if v.Value != "" {
					if total > 0 {
						conditions += " AND"
					}
					conditions += filterByOperator(v.Field, v.Operator, v.Value)
					total++
				}
			}
			conditions += ")"
		}

		var orders = ""
		if len(request.Orders) > 0 {
			order_total := 0

			for _, v := range request.Orders {
				if order_total > 0 {
					orders += ", "
				}
				orders += fmt.Sprintf("%v %v", v.Field, v.Direction)
				order_total++
			}
		} else {
			orders = "COALESCE(submitted_at, created_at) DESC"
		}

		if err := r.DB.
			//Select("ROW_NUMBER() OVER( ORDER BY COALESCE(submitted_at, created_at) DESC)").
			// Debug().
			Order(orders).
			Select(strings.Join(selectedFields, ", ")).Table(fmt.Sprintf("vw_%v", tableName)).
			// Where(canRead).Where(conditions).Where(searchConditions).
			Where(conditions).
			Offset(offset).Limit(pageSize).Find(&records).Error; err != nil {
			return nil, err
		}
	}

	response.RecordsFiltered = int(len(records))
	response.Data = records
	return response, nil
}