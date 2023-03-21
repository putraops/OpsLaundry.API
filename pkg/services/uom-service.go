package services

import (
	"opslaundry/pkg/commons"
	"opslaundry/pkg/models"
	"opslaundry/pkg/models/views"
	"opslaundry/pkg/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UomService interface {
	Create(context *gin.Context, record models.Uom) (interface{}, error)
	Update(context *gin.Context, record models.Uom) (interface{}, error)
	GetPagination(context *gin.Context, request commons.DataTableRequest) interface{}
	GetAll(context *gin.Context) (interface{}, error)
	GetById(context *gin.Context, id string) (interface{}, error)
	DeleteById(context *gin.Context, id string) error
}

type uomService struct {
	uomRepository  repository.UomRepository
	baseRepository repository.BaseRepository
}

func NewUomService(db *gorm.DB) UomService {
	return &uomService{
		uomRepository:  repository.NewUomRepository(db),
		baseRepository: repository.NewBaseRepository(db),
	}
}

func (r *uomService) Create(context *gin.Context, record models.Uom) (interface{}, error) {
	return r.uomRepository.Create(context, record)
}

func (r *uomService) Update(context *gin.Context, record models.Uom) (interface{}, error) {
	return r.uomRepository.Update(context, record)
}

func (r *uomService) GetPagination(context *gin.Context, request commons.DataTableRequest) interface{} {
	context.Set("table", models.Uom{})
	context.Set("table_name", models.Uom{}.TableName())
	context.Set("view", views.Uom{})
	return r.baseRepository.GetPagination(context, request, false)
}

func (r *uomService) GetAll(context *gin.Context) (interface{}, error) {
	return r.uomRepository.GetAll(context)
}

func (r *uomService) GetById(context *gin.Context, id string) (interface{}, error) {
	return r.uomRepository.GetById(context, id)
}

func (r *uomService) DeleteById(context *gin.Context, id string) error {
	return r.uomRepository.DeleteById(context, id)
}
