package services

import (
	"opslaundry/pkg/commons"
	"opslaundry/pkg/models"
	"opslaundry/pkg/models/views"
	"opslaundry/pkg/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TenantService interface {
	Create(context *gin.Context, record models.Tenant) (interface{}, error)
	Update(context *gin.Context, record models.Tenant) (interface{}, error)
	GetPagination(context *gin.Context, request commons.DataTableRequest) (interface{}, error)
	GetAll(context *gin.Context) (interface{}, error)
	GetById(context *gin.Context, id string) (interface{}, error)
	DeleteById(context *gin.Context, id string) error
}

type tenantService struct {
	tenantRepository repository.TenantRepository
	baseRepository   repository.BaseRepository
}

func NewTenantService(db *gorm.DB) TenantService {
	return &tenantService{
		tenantRepository: repository.NewTenantRepository(db),
		baseRepository:   repository.NewBaseRepository(db),
	}
}

func (r *tenantService) Create(context *gin.Context, record models.Tenant) (interface{}, error) {
	return r.tenantRepository.Create(context, record)
}

func (r *tenantService) Update(context *gin.Context, record models.Tenant) (interface{}, error) {
	return r.tenantRepository.Update(context, record)
}

func (r tenantService) GetPagination(context *gin.Context, request commons.DataTableRequest) (interface{}, error) {
	context.Set("table", models.Tenant{})
	context.Set("table_name", models.Tenant{}.TableName())
	context.Set("view", views.Tenant{})
	return r.baseRepository.GetPagination(context, request, false)
}

func (r *tenantService) GetAll(context *gin.Context) (interface{}, error) {
	return r.tenantRepository.GetAll(context)
}

func (r *tenantService) GetById(context *gin.Context, id string) (interface{}, error) {
	return r.tenantRepository.GetById(context, id)
}

func (r *tenantService) DeleteById(context *gin.Context, id string) error {
	return r.tenantRepository.DeleteById(context, id)
}
