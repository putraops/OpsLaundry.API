package repository

import (
	"encoding/json"
	"opslaundry/pkg/models"
	"opslaundry/pkg/models/views"
	"opslaundry/pkg/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TenantRepository interface {
	Create(context *gin.Context, record models.Tenant) (interface{}, error)
	Update(context *gin.Context, record models.Tenant) (interface{}, error)
	GetAll(context *gin.Context) (interface{}, error)
	GetById(context *gin.Context, id string) (interface{}, error)
	DeleteById(context *gin.Context, id string) error
}

type tenantRepository struct {
	DB             *gorm.DB
	baseRepository BaseRepository
}

func NewTenantRepository(db *gorm.DB) TenantRepository {
	return &tenantRepository{
		DB:             db,
		baseRepository: NewBaseRepository(db),
	}
}

func (r tenantRepository) Create(context *gin.Context, record models.Tenant) (interface{}, error) {
	if err := r.DB.WithContext(context.Request.Context()).Create(&record).Error; err != nil {
		return nil, err
	}

	return record, nil
}

func (r tenantRepository) Update(context *gin.Context, record models.Tenant) (interface{}, error) {
	var obj models.Tenant

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

func (r tenantRepository) GetAll(context *gin.Context) (interface{}, error) {
	var records []views.Tenant
	if err := r.DB.Find(&records).Error; err != nil {
		return nil, err
	}

	return records, nil
}

func (r tenantRepository) GetById(context *gin.Context, id string) (interface{}, error) {
	var record views.Tenant
	if err := r.DB.Where("id = ?", id).Take(&record).Error; err != nil {
		return nil, err
	}

	return record, nil
}

func (r tenantRepository) DeleteById(context *gin.Context, id string) error {
	var record models.Tenant
	if err := r.DB.Where("id = ?", id).First(&record).Error; err != nil {
		return err
	}

	if err := r.DB.WithContext(context.Request.Context()).Delete(&record).Error; err != nil {
		return err
	}

	return nil
}
