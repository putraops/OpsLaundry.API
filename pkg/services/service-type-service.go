package services

import (
	"opslaundry/pkg/commons"
	"opslaundry/pkg/models"
	"opslaundry/pkg/models/views"
	"opslaundry/pkg/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ServiceTypeService interface {
	Create(context *gin.Context, record models.ServiceType) (interface{}, error)
	Update(context *gin.Context, record models.ServiceType) (interface{}, error)
	GetPagination(context *gin.Context, request commons.DataTableRequest) interface{}
	GetAll(context *gin.Context) (interface{}, error)
	GetById(context *gin.Context, id string) (interface{}, error)
	DeleteById(context *gin.Context, id string) error
}

type serviceTypeService struct {
	serviceTypeRepository repository.ServiceTypeRepository
	baseRepository        repository.BaseRepository
}

func NewServiceTypeService(db *gorm.DB) ServiceTypeService {
	return &serviceTypeService{
		serviceTypeRepository: repository.NewServiceTypeRepository(db),
		baseRepository:        repository.NewBaseRepository(db),
	}
}

func (r *serviceTypeService) Create(context *gin.Context, record models.ServiceType) (interface{}, error) {
	return r.serviceTypeRepository.Create(context, record)
}

func (r *serviceTypeService) Update(context *gin.Context, record models.ServiceType) (interface{}, error) {
	return r.serviceTypeRepository.Update(context, record)
}

func (r *serviceTypeService) GetPagination(context *gin.Context, request commons.DataTableRequest) interface{} {
	context.Set("table", models.ServiceType{})
	context.Set("table_name", models.ServiceType{}.TableName())
	context.Set("view", views.ServiceType{})
	return r.baseRepository.GetPagination(context, request, false)
}

func (r *serviceTypeService) GetAll(context *gin.Context) (interface{}, error) {
	return r.serviceTypeRepository.GetAll(context)
}

func (r *serviceTypeService) GetById(context *gin.Context, id string) (interface{}, error) {
	return r.serviceTypeRepository.GetById(context, id)
}

func (r *serviceTypeService) DeleteById(context *gin.Context, id string) error {
	return r.serviceTypeRepository.DeleteById(context, id)
}
