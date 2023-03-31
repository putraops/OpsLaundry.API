package controllers

import (
	"errors"
	"net/http"
	"opslaundry/pkg/commons"
	"opslaundry/pkg/constants"
	"opslaundry/pkg/models"
	"opslaundry/pkg/services"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ServiceTypeController interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	GetPagination(context *gin.Context)
	GetAll(context *gin.Context)
	GetById(context *gin.Context)
	DeleteById(context *gin.Context)
}

type serviceTypeController struct {
	serviceTypeService services.ServiceTypeService
	logger             commons.OpsLogger
}

func NewServiceTypeController(db *gorm.DB) ServiceTypeController {
	return &serviceTypeController{
		serviceTypeService: services.NewServiceTypeService(db),
		logger:             commons.NewLogger(),
	}
}

// @Tags         Service Type
// @Security 	 BearerAuth
// @Summary
// @Description
// @Accept       json
// @Produce      json
// @Param        Request body models.ServiceType true "ServiceType"
// @Failure 	 200,204,400,401,404 {object} object
// @Router       /service_type/create [post]
func (r *serviceTypeController) Create(context *gin.Context) {
	var record models.ServiceType

	if err := context.BindJSON(&record); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, commons.ErrorResponse(err.Error()))
		return
	}

	result, err := r.serviceTypeService.Create(context, record)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, commons.ErrorResponse(err.Error()))
		return
	}

	context.JSON(http.StatusOK, result)
}

// @Tags         Service Type
// @Security 	 BearerAuth
// @Summary
// @Description
// @Accept       json
// @Produce      json
// @Param        Request body models.ServiceType true "serviceType"
// @Failure 	 200,204,400,401,404 {object} object
// @Router       /service_type/update [patch]
func (r *serviceTypeController) Update(context *gin.Context) {
	var record models.ServiceType

	if err := context.BindJSON(&record); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, commons.ErrorResponse(err.Error()))
		return
	}

	result, err := r.serviceTypeService.Update(context, record)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.AbortWithStatus(http.StatusNoContent)
			return
		}
		context.AbortWithStatusJSON(http.StatusBadRequest, commons.ErrorResponse(err.Error()))
		return
	}

	context.JSON(http.StatusOK, result)
}

// @Tags         Service Type
// @Security 	 BearerAuth
// @Accept       json
// @Produce      json
// @Param        body body commons.DataTableRequest true " "
// @Failure 	 200,204,400,401,404 {object} object
// @Router       /service_type/getPagination [post]
func (r *serviceTypeController) GetPagination(context *gin.Context) {
	var req commons.DataTableRequest

	if err := context.BindJSON(&req); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, commons.ErrorResponse(err.Error()))
		return
	}

	result := r.serviceTypeService.GetPagination(context, req).(commons.DataTableResponse)
	context.JSON(http.StatusOK, result)
}

// @Tags         Service Type
// @Security 	 BearerAuth
// @Summary
// @Description
// @Accept       json
// @Produce      json
// @Failure 	 200,204,400,401,404 {object} object
// @Router       /service_type/getAll [get]
func (r *serviceTypeController) GetAll(context *gin.Context) {
	result, err := r.serviceTypeService.GetAll(context)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.AbortWithStatus(http.StatusNoContent)
			return
		}
		context.AbortWithStatusJSON(http.StatusBadRequest, commons.ErrorResponse(err.Error()))
		return
	}

	context.JSON(http.StatusOK, result)
}

// @Tags         Service Type
// @Security 	 BearerAuth
// @Summary
// @Description
// @Accept       json
// @Produce      json
// @Param        id path string true "id"
// @Failure 	 200,204,400,401,404 {object} object
// @Router       /service_type/getById/{id} [get]
func (r *serviceTypeController) GetById(context *gin.Context) {
	id := context.Param("id")
	if strings.Trim(id, " ") == "" {
		context.AbortWithStatusJSON(http.StatusBadRequest, commons.ErrorResponse(constants.NoIdFound))
		return
	}

	result, err := r.serviceTypeService.GetById(context, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.AbortWithStatus(http.StatusNoContent)
			return
		}
		context.AbortWithStatusJSON(http.StatusBadRequest, commons.ErrorResponse(err.Error()))
		return
	}

	context.JSON(http.StatusOK, result)
}

// @Tags         Service Type
// @Security 	 BearerAuth
// @Summary
// @Description
// @Accept       json
// @Produce      json
// @Param        id path string true "id"
// @Success      200 {string} string
// @Failure 	 200,204,400,401,404 {object} object
// @Router       /service_type/deleteById/{id} [delete]
func (r *serviceTypeController) DeleteById(context *gin.Context) {
	id := context.Param("id")
	if strings.Trim(id, " ") == "" {
		context.AbortWithStatusJSON(http.StatusBadRequest, commons.ErrorResponse(constants.NoIdFound))
		return
	}

	if err := r.serviceTypeService.DeleteById(context, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.AbortWithStatus(http.StatusNoContent)
			return
		}
		context.AbortWithStatusJSON(http.StatusBadRequest, commons.ErrorResponse(err.Error()))
		return
	}

	context.JSON(http.StatusOK, constants.DeletedMessage)
}
