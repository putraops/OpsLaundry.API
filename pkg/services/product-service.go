package services

import (
	"opslaundry/pkg/commons"
	"opslaundry/pkg/models"
	"opslaundry/pkg/models/views"
	"opslaundry/pkg/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductService interface {
	Create(context *gin.Context, record models.Product) (interface{}, error)
	Update(context *gin.Context, record models.Product) (interface{}, error)
	GetPagination(context *gin.Context, request commons.DataTableRequest) (interface{}, error)
	GetAll(context *gin.Context) (interface{}, error)
	GetById(context *gin.Context, id string) (interface{}, error)
	DeleteById(context *gin.Context, id string) error
}

type productService struct {
	productRepository repository.ProductRepository
	baseRepository    repository.BaseRepository
}

func NewProductService(db *gorm.DB) ProductService {
	return &productService{
		productRepository: repository.NewProductRepository(db),
		baseRepository:    repository.NewBaseRepository(db),
	}
}

func (r *productService) Create(context *gin.Context, record models.Product) (interface{}, error) {
	return r.productRepository.Create(context, record)
}

func (r *productService) Update(context *gin.Context, record models.Product) (interface{}, error) {
	return r.productRepository.Update(context, record)
}

func (r *productService) GetPagination(context *gin.Context, request commons.DataTableRequest) (interface{}, error) {
	context.Set("table", models.Product{})
	context.Set("table_name", models.Product{}.TableName())
	context.Set("view", views.Product{})
	return r.baseRepository.GetPagination(context, request, false)
}

func (r *productService) GetAll(context *gin.Context) (interface{}, error) {
	return r.productRepository.GetAll(context)
}

func (r *productService) GetById(context *gin.Context, id string) (interface{}, error) {
	return r.productRepository.GetById(context, id)
}

func (r *productService) DeleteById(context *gin.Context, id string) error {
	return r.productRepository.DeleteById(context, id)
}
