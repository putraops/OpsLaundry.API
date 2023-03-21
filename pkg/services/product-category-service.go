package services

import (
	"opslaundry/pkg/commons"
	"opslaundry/pkg/models"
	"opslaundry/pkg/models/views"
	"opslaundry/pkg/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductCategoryService interface {
	Create(context *gin.Context, record models.ProductCategory) (interface{}, error)
	Update(context *gin.Context, record models.ProductCategory) (interface{}, error)
	GetPagination(context *gin.Context, request commons.DataTableRequest) interface{}
	GetAll(context *gin.Context) (interface{}, error)
	GetById(context *gin.Context, id string) (interface{}, error)
	DeleteById(context *gin.Context, id string) error
}

type productCategoryService struct {
	productCategoryRepository repository.ProductCategoryRepository
	baseRepository            repository.BaseRepository
}

func NewProductCategoryService(db *gorm.DB) ProductCategoryService {
	return &productCategoryService{
		productCategoryRepository: repository.NewProductCategoryRepository(db),
		baseRepository:            repository.NewBaseRepository(db),
	}
}

func (r *productCategoryService) Create(context *gin.Context, record models.ProductCategory) (interface{}, error) {
	return r.productCategoryRepository.Create(context, record)
}

func (r *productCategoryService) Update(context *gin.Context, record models.ProductCategory) (interface{}, error) {
	return r.productCategoryRepository.Update(context, record)
}

func (r *productCategoryService) GetPagination(context *gin.Context, request commons.DataTableRequest) interface{} {
	context.Set("table", models.ProductCategory{})
	context.Set("table_name", models.ProductCategory{}.TableName())
	context.Set("view", views.ProductCategory{})
	return r.baseRepository.GetPagination(context, request, false)
}

func (r *productCategoryService) GetAll(context *gin.Context) (interface{}, error) {
	return r.productCategoryRepository.GetAll(context)
}

func (r *productCategoryService) GetById(context *gin.Context, id string) (interface{}, error) {
	return r.productCategoryRepository.GetById(context, id)
}

func (r *productCategoryService) DeleteById(context *gin.Context, id string) error {
	return r.productCategoryRepository.DeleteById(context, id)
}
