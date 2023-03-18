package services

import (
	"errors"
	"opslaundry/pkg/constants"
	"opslaundry/pkg/models"
	"opslaundry/pkg/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// -- This is user Organization
type OrganizationService interface {
	Update(context *gin.Context, record models.Organization) (interface{}, error)
	// Save(context *gin.Context, record models.Organization) helper.StandartResult
	// Lookup(context *gin.Context, request helper.ReactSelectRequest) helper.StandartResult
	// GetPagination(context *gin.Context, request commons.DataTableRequest) interface{}
	// GetAll(context *gin.Context, filter map[string]interface{}) helper.StandartResult
	// GetById(context *gin.Context, id string) helper.StandartResult
	// DeleteById(context *gin.Context, id string) helper.StandartResult
}

type organizationService struct {
	jwtService             JWTService
	organizationRepository repository.OrganizationRepository
}

func NewOrganizationService(db *gorm.DB, jwtService JWTService) OrganizationService {
	return &organizationService{
		jwtService:             jwtService,
		organizationRepository: repository.NewOrganizationRepository(db),
	}
}

func (r organizationService) Update(context *gin.Context, record models.Organization) (interface{}, error) {
	result, err := r.organizationRepository.Update(context, record)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(constants.NoDataFound)
		}
		return nil, err
	}
	return result, nil
}

// func (r organizationService) Lookup(context *gin.Context, request helper.ReactSelectRequest) helper.StandartResult {
// 	return r.organizationRepository.Lookup(context, request)
// }

// func (r organizationService) GetPagination(context *gin.Context, request commons.DataTableRequest) interface{} {
// 	return r.organizationRepository.GetPagination(context, request)
// }

// func (r organizationService) GetAll(context *gin.Context, filter map[string]interface{}) helper.StandartResult {
// 	return r.organizationRepository.GetAll(context, filter)
// }

// func (r organizationService) GetById(context *gin.Context, recordId string) helper.StandartResult {
// 	return r.organizationRepository.GetById(context, recordId)
// }

// func (r organizationService) DeleteById(context *gin.Context, id string) helper.StandartResult {
// 	return r.organizationRepository.DeleteById(context, id)
// }
