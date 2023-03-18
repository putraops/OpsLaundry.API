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

type TenantController interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	GetPagination(context *gin.Context)
	GetAll(context *gin.Context)
	GetById(context *gin.Context)
	DeleteById(context *gin.Context)
}

type tenantController struct {
	tenantService services.TenantService
	logger        commons.OpsLogger
}

func NewTenantController(db *gorm.DB) TenantController {
	return &tenantController{
		tenantService: services.NewTenantService(db),
		logger:        commons.NewLogger(),
	}
}

// @Tags         Tenant
// @Security 	 BearerAuth
// @Summary
// @Description
// @Accept       json
// @Produce      json
// @Param        Request body models.Tenant true "tenant"
// @Failure 	 200,204,400,401,404 {object} object
// @Router       /tenant/create [post]
func (r tenantController) Create(context *gin.Context) {
	var record models.Tenant

	if err := context.BindJSON(&record); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, commons.ErrorResponse(err.Error()))
		return
	}

	result, err := r.tenantService.Create(context, record)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, commons.ErrorResponse(err.Error()))
		return
	}

	context.JSON(http.StatusOK, result)
	return
}

// @Tags         Tenant
// @Security 	 BearerAuth
// @Summary
// @Description
// @Accept       json
// @Produce      json
// @Param        Request body models.Tenant true "tenant"
// @Failure 	 200,204,400,401,404 {object} object
// @Router       /tenant/update [patch]
func (r tenantController) Update(context *gin.Context) {
	var record models.Tenant

	if err := context.BindJSON(&record); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, commons.ErrorResponse(err.Error()))
		return
	}

	result, err := r.tenantService.Update(context, record)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.AbortWithStatus(http.StatusNoContent)
			return
		}
		context.AbortWithStatusJSON(http.StatusBadRequest, commons.ErrorResponse(err.Error()))
		return
	}

	context.JSON(http.StatusOK, result)
	return
}

// @Tags         Tenant
// @Security 	 BearerAuth
// @Accept       json
// @Produce      json
// @Param        body body commons.DataTableRequest true " "
// @Failure 	 200,204,400,401,404 {object} object
// @Router       /tenant/getPagination [post]
func (r tenantController) GetPagination(context *gin.Context) {
	var req commons.DataTableRequest

	if err := context.BindJSON(&req); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, commons.ErrorResponse(err.Error()))
		return
	}

	result, err := r.tenantService.GetPagination(context, req)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, commons.ErrorResponse(err.Error()))
		return
	}
	context.JSON(http.StatusOK, result)
	return
}

// @Tags         Tenant
// @Security 	 BearerAuth
// @Summary
// @Description
// @Accept       json
// @Produce      json
// @Failure 	 200,204,400,401,404 {object} object
// @Router       /tenant/getAll [get]
func (r tenantController) GetAll(context *gin.Context) {
	result, err := r.tenantService.GetAll(context)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.AbortWithStatus(http.StatusNoContent)
			return
		}
		context.AbortWithStatusJSON(http.StatusBadRequest, commons.ErrorResponse(err.Error()))
		return
	}

	context.JSON(http.StatusOK, result)
	return
}

// @Tags         Tenant
// @Security 	 BearerAuth
// @Summary
// @Description
// @Accept       json
// @Produce      json
// @Param        id path string true "id"
// @Failure 	 200,204,400,401,404 {object} object
// @Router       /tenant/getById/{id} [get]
func (r tenantController) GetById(context *gin.Context) {
	id := context.Param("id")
	if strings.Trim(id, " ") == "" {
		context.AbortWithStatusJSON(http.StatusBadRequest, commons.ErrorResponse(constants.NoIdFound))
		return
	}

	result, err := r.tenantService.GetById(context, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.AbortWithStatus(http.StatusNoContent)
			return
		}
		context.AbortWithStatusJSON(http.StatusBadRequest, commons.ErrorResponse(err.Error()))
		return
	}

	context.JSON(http.StatusOK, result)
	return
}

// @Tags         Tenant
// @Security 	 BearerAuth
// @Summary
// @Description
// @Accept       json
// @Produce      json
// @Param        id path string true "id"
// @Success      200 {string} string
// @Failure 	 200,204,400,401,404 {object} object
// @Router       /tenant/deleteById/{id} [delete]
func (r tenantController) DeleteById(context *gin.Context) {
	id := context.Param("id")
	if strings.Trim(id, " ") == "" {
		context.AbortWithStatusJSON(http.StatusBadRequest, commons.ErrorResponse(constants.NoIdFound))
		return
	}

	if err := r.tenantService.DeleteById(context, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.AbortWithStatus(http.StatusNoContent)
			return
		}
		context.AbortWithStatusJSON(http.StatusBadRequest, commons.ErrorResponse(err.Error()))
		return
	}

	context.JSON(http.StatusOK, constants.DeletedMessage)
	return
}
