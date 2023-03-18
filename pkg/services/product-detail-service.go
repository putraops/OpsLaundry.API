package services

import (
	"opslaundry/pkg/commons"
	"opslaundry/pkg/models"
	"opslaundry/pkg/models/views"
	"opslaundry/pkg/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductDetailService interface {
	Create(context *gin.Context, record models.ProductDetail) (interface{}, error)
	Update(context *gin.Context, record models.ProductDetail) (interface{}, error)
	GetPagination(context *gin.Context, request commons.DataTableRequest) (interface{}, error)
	GetAll(context *gin.Context) (interface{}, error)
	GetById(context *gin.Context, id string) (interface{}, error)
	DeleteById(context *gin.Context, id string) error
}

type productDetailService struct {
	productDetailRepository repository.ProductDetailRepository
	baseRepository          repository.BaseRepository
}

func NewProductDetailService(db *gorm.DB) ProductDetailService {
	return &productDetailService{
		productDetailRepository: repository.NewProductDetailRepository(db),
		baseRepository:          repository.NewBaseRepository(db),
	}
}

func (r *productDetailService) Create(context *gin.Context, record models.ProductDetail) (interface{}, error) {
	return r.productDetailRepository.Create(context, record)
}

func (r *productDetailService) Update(context *gin.Context, record models.ProductDetail) (interface{}, error) {
	return r.productDetailRepository.Update(context, record)
}

func (r *productDetailService) GetPagination(context *gin.Context, request commons.DataTableRequest) (interface{}, error) {
	context.Set("table", models.ProductDetail{})
	context.Set("table_name", models.ProductDetail{}.TableName())
	context.Set("view", views.ProductDetail{})
	return r.baseRepository.GetPagination(context, request, false)
}

func (r *productDetailService) GetAll(context *gin.Context) (interface{}, error) {
	return r.productDetailRepository.GetAll(context)
}

func (r *productDetailService) GetById(context *gin.Context, id string) (interface{}, error) {
	return r.productDetailRepository.GetById(context, id)
}

func (r *productDetailService) DeleteById(context *gin.Context, id string) error {
	return r.productDetailRepository.DeleteById(context, id)
}
