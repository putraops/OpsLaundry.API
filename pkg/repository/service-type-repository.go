package repository

import (
	"encoding/json"
	"opslaundry/pkg/models"
	"opslaundry/pkg/models/views"
	"opslaundry/pkg/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ServiceTypeRepository interface {
	Create(context *gin.Context, record models.ServiceType) (interface{}, error)
	Update(context *gin.Context, record models.ServiceType) (interface{}, error)
	GetAll(context *gin.Context) (interface{}, error)
	GetById(context *gin.Context, id string) (interface{}, error)
	DeleteById(context *gin.Context, id string) error
}

type serviceTypeRepository struct {
	DB             *gorm.DB
	baseRepository BaseRepository
}

func NewServiceTypeRepository(db *gorm.DB) ServiceTypeRepository {
	return &serviceTypeRepository{
		DB:             db,
		baseRepository: NewBaseRepository(db),
	}
}

func (r *serviceTypeRepository) Create(context *gin.Context, record models.ServiceType) (interface{}, error) {
	if err := r.DB.WithContext(context.Request.Context()).Create(&record).Error; err != nil {
		return nil, err
	}

	return record, nil
}

func (r *serviceTypeRepository) Update(context *gin.Context, record models.ServiceType) (interface{}, error) {
	var obj models.ServiceType

	if err := r.DB.First(&obj, "id = ?", record.Id).Error; err != nil {
		return nil, err
	}

	mapResult, err := utils.MapFields(obj, record)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(mapResult, &obj); err != nil {
		return nil, err
	}

	if err := r.DB.WithContext(context.Request.Context()).Save(&obj).Error; err != nil {
		return nil, err
	}

	return obj, nil
}

func (r *serviceTypeRepository) GetAll(context *gin.Context) (interface{}, error) {
	var records []views.ServiceType
	if err := r.DB.Find(&records).Error; err != nil {
		return nil, err
	}

	return records, nil
}

func (r *serviceTypeRepository) GetById(context *gin.Context, id string) (interface{}, error) {
	var record views.ServiceType
	if err := r.DB.Where("id = ?", id).Take(&record).Error; err != nil {
		return nil, err
	}

	return record, nil
}

func (r *serviceTypeRepository) DeleteById(context *gin.Context, id string) error {
	var record models.ServiceType
	if err := r.DB.Where("id = ?", id).First(&record).Error; err != nil {
		return err
	}

	if err := r.DB.WithContext(context.Request.Context()).Delete(&record).Error; err != nil {
		return err
	}

	return nil
}
