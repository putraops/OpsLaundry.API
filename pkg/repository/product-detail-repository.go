package repository

import (
	"encoding/json"
	"opslaundry/pkg/models"
	"opslaundry/pkg/models/views"
	"opslaundry/pkg/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductDetailRepository interface {
	Create(context *gin.Context, record models.ProductDetail) (interface{}, error)
	Update(context *gin.Context, record models.ProductDetail) (interface{}, error)
	GetAll(context *gin.Context) (interface{}, error)
	GetById(context *gin.Context, id string) (interface{}, error)
	DeleteById(context *gin.Context, id string) error
}

type productDetailRepository struct {
	DB             *gorm.DB
	baseRepository BaseRepository
}

func NewProductDetailRepository(db *gorm.DB) ProductDetailRepository {
	return &productDetailRepository{
		DB:             db,
		baseRepository: NewBaseRepository(db),
	}
}

func (r *productDetailRepository) Create(context *gin.Context, record models.ProductDetail) (interface{}, error) {
	if err := r.DB.WithContext(context.Request.Context()).Create(&record).Error; err != nil {
		return nil, err
	}

	return record, nil
}

func (r *productDetailRepository) Update(context *gin.Context, record models.ProductDetail) (interface{}, error) {
	var obj models.ProductDetail

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

func (r *productDetailRepository) GetAll(context *gin.Context) (interface{}, error) {
	var records []views.ProductDetail
	if err := r.DB.Find(&records).Error; err != nil {
		return nil, err
	}

	return records, nil
}

func (r *productDetailRepository) GetById(context *gin.Context, id string) (interface{}, error) {
	var record views.ProductDetail
	if err := r.DB.Where("id = ?", id).Take(&record).Error; err != nil {
		return nil, err
	}

	return record, nil
}

func (r *productDetailRepository) DeleteById(context *gin.Context, id string) error {
	var record models.ProductDetail
	if err := r.DB.Where("id = ?", id).First(&record).Error; err != nil {
		return err
	}

	if err := r.DB.WithContext(context.Request.Context()).Delete(&record).Error; err != nil {
		return err
	}

	return nil
}
