package repository

import (
	"encoding/json"
	"opslaundry/pkg/models"
	"opslaundry/pkg/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrganizationRepository interface {
	// CreateOrUpdate(context *gin.Context, record models.Organization) helper.StandartResult
	Update(context *gin.Context, record models.Organization) (interface{}, error)
	// Lookup(context *gin.Context, request helper.ReactSelectRequest) helper.StandartResult
	// GetPagination(context *gin.Context, request commons.DataTableRequest) interface{}
	// GetAll(context *gin.Context, filter map[string]interface{}) helper.StandartResult
	// GetById(context *gin.Context, id string) helper.StandartResult

	// DeleteById(context *gin.Context, id string) helper.StandartResult
}

type organizationRepository struct {
	DB             *gorm.DB
	baseRepository BaseRepository
}

func NewOrganizationRepository(db *gorm.DB) OrganizationRepository {
	return &organizationRepository{
		DB:             db,
		baseRepository: NewBaseRepository(db),
	}
}

func (r organizationRepository) Update(context *gin.Context, record models.Organization) (interface{}, error) {
	// var userIdentity = getUser(context)

	// _context := r.DB.Set("identity", nil)

	var oldRecord models.Organization
	if err := r.DB.First(&oldRecord, record.Id).Error; err != nil {
		return nil, err
	}

	mapResult, err := utils.MapFields(oldRecord, record)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(mapResult, &oldRecord); err != nil {
		return nil, err
	}

	if err := r.DB.Save(&oldRecord).Error; err != nil {
		return nil, err
	}
	return oldRecord, nil
}

// func (r organizationRepository) CreateOrUpdate(context *gin.Context, record models.Organization) helper.StandartResult {
// 	var result helper.StandartResult
// 	result.Success = true

// 	var userIdentity = getUser(context)
// 	_context := r.DB.Set("identity", userIdentity)

// 	if record.Id == "" {
// 		fmt.Println("Creating New Record....")
// 		result.Message = helper.NewDataMessage
// 		_context.Transaction(func(tx *gorm.DB) error {
// 			if errTx := tx.Create(&record).Error; errTx != nil {
// 				result.Success = false
// 				result.Message = fmt.Sprintf("%v", errTx)
// 				if strings.Contains(result.Message, "violates unique constraint") {
// 					result.Message = helper.ViolatesUniqueConstraintMessage
// 				}
// 				return errTx
// 			} else {
// 				result.Data = record
// 			}
// 			return nil
// 		})

// 	} else {
// 		fmt.Println("Updating Record....")
// 		var oldRecord models.Organization
// 		r.DB.Where("id = ?", record.Id).Where(filteredByIdentity(userIdentity, true, false)).First(&oldRecord)
// 		if oldRecord.Id == "" {
// 			result.Success = false
// 			result.Message = helper.NoDataFound
// 			return result
// 		}

// 		mapResult := r.baseRepository.MapFields(oldRecord, record)
// 		if !mapResult.Success {
// 			fmt.Println("Failed to MapFields")
// 			return mapResult
// 		}
// 		json.Unmarshal(mapResult.Data.([]byte), &oldRecord)

// 		result.Message = helper.UpdateDataMessage
// 		_context.Transaction(func(tx *gorm.DB) error {
// 			if errTx := tx.Save(&oldRecord).Error; errTx != nil {
// 				result.Success = false
// 				result.Message = fmt.Sprintf("%v", errTx)
// 				return errTx
// 			} else {
// 				result.Data = oldRecord
// 			}
// 			return nil
// 		})
// 	}
// 	return result
// }

// func (r organizationRepository) Lookup(context *gin.Context, request helper.ReactSelectRequest) helper.StandartResult {
// 	return r.baseRepository.Lookup(context, request, false)
// }

// func (r organizationRepository) GetPagination(context *gin.Context, request commons.DataTableRequest) interface{} {
// 	return r.baseRepository.GetPagination(context, request, false)
// }

// func (r organizationRepository) GetAll(context *gin.Context, filter map[string]interface{}) helper.StandartResult {
// 	var userIdentity = getUser(context)
// 	var result helper.StandartResult
// 	result.Success = true
// 	result.Message = helper.OkMessage

// 	var records []entity_view_models.EntityOrganizationView
// 	r.DB.Debug().Where(filteredByIdentity(userIdentity, true, false)).Where(filter).Find(&records)

// 	if len(records) <= 0 {
// 		result.Success = false
// 		result.Message = helper.NoDataFound
// 	} else {
// 		result.Data = records
// 	}
// 	return result
// }

// func (r organizationRepository) DeleteById(context *gin.Context, id string) helper.StandartResult {
// 	var userIdentity = getUser(context)

// 	var record models.Organization
// 	r.DB.Where("id = ?", id).Where(filteredByIdentity(userIdentity, true, false)).First(&record)
// 	if record.Id == "" {
// 		return helper.StandartResult{Success: false, Message: helper.NoDataFound, Data: nil}
// 	}

// 	_context := r.DB.Set("identity", userIdentity)
// 	_context.Transaction(func(tx *gorm.DB) error {
// 		if errTx := tx.Delete(&record).Error; errTx != nil {
// 			// return any error will rollback
// 			return errTx
// 		}
// 		return nil
// 	})
// 	return helper.StandartResult{Success: true, Message: helper.OkMessage, Data: nil}
// }

// func (r organizationRepository) GetById(context *gin.Context, id string) helper.StandartResult {
// 	var userIdentity = getUser(context)
// 	var record entity_view_models.EntityOrganizationView

// 	r.DB.Where("id = ?", id).Where(filteredByIdentity(userIdentity, true, false)).First(&record)
// 	if record.Id == "" {
// 		res := helper.StandartResult{Success: false, Message: helper.NoDataFound, Data: nil}
// 		return res
// 	}

// 	res := helper.StandartResult{Success: true, Message: "Ok", Data: record}
// 	return res
// }
