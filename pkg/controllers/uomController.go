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

type UomController interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	GetPagination(context *gin.Context)
	GetAll(context *gin.Context)
	GetById(context *gin.Context)
	DeleteById(context *gin.Context)
}

type uomController struct {
	uomService services.UomService
	logger     commons.OpsLogger
}

func NewUomController(db *gorm.DB) UomController {
	return &uomController{
		uomService: services.NewUomService(db),
		logger:     commons.NewLogger(),
	}
}

// @Tags         Unit of Measurement
// @Security 	 BearerAuth
// @Summary
// @Description
// @Accept       json
// @Produce      json
// @Param        Request body models.Uom true "Uom"
// @Failure 	 200,204,400,401,404 {object} object
// @Router       /uom/create [post]
func (r *uomController) Create(context *gin.Context) {
	var record models.Uom

	if err := context.BindJSON(&record); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, commons.ErrorResponse(err.Error()))
		return
	}

	result, err := r.uomService.Create(context, record)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, commons.ErrorResponse(err.Error()))
		return
	}

	context.JSON(http.StatusOK, result)
	return
}

// @Tags         Unit of Measurement
// @Security 	 BearerAuth
// @Summary
// @Description
// @Accept       json
// @Produce      json
// @Param        Request body models.Uom true "uom"
// @Failure 	 200,204,400,401,404 {object} object
// @Router       /uom/update [patch]
func (r *uomController) Update(context *gin.Context) {
	var record models.Uom

	if err := context.BindJSON(&record); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, commons.ErrorResponse(err.Error()))
		return
	}

	result, err := r.uomService.Update(context, record)
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

// @Tags         Unit of Measurement
// @Security 	 BearerAuth
// @Accept       json
// @Produce      json
// @Param        body body commons.DataTableRequest true " "
// @Failure 	 200,204,400,401,404 {object} object
// @Router       /uom/getPagination [post]
func (r *uomController) GetPagination(context *gin.Context) {
	var req commons.DataTableRequest

	if err := context.BindJSON(&req); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, commons.ErrorResponse(err.Error()))
		return
	}

	result := r.uomService.GetPagination(context, req).(commons.DataTableResponse)
	context.JSON(http.StatusOK, result)
	return
}

// @Tags         Unit of Measurement
// @Security 	 BearerAuth
// @Summary
// @Description
// @Accept       json
// @Produce      json
// @Failure 	 200,204,400,401,404 {object} object
// @Router       /uom/getAll [get]
func (r *uomController) GetAll(context *gin.Context) {
	result, err := r.uomService.GetAll(context)

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

// @Tags         Unit of Measurement
// @Security 	 BearerAuth
// @Summary
// @Description
// @Accept       json
// @Produce      json
// @Param        id path string true "id"
// @Failure 	 200,204,400,401,404 {object} object
// @Router       /uom/getById/{id} [get]
func (r *uomController) GetById(context *gin.Context) {
	id := context.Param("id")
	if strings.Trim(id, " ") == "" {
		context.AbortWithStatusJSON(http.StatusBadRequest, commons.ErrorResponse(constants.NoIdFound))
		return
	}

	result, err := r.uomService.GetById(context, id)
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

// @Tags         Unit of Measurement
// @Security 	 BearerAuth
// @Summary
// @Description
// @Accept       json
// @Produce      json
// @Param        id path string true "id"
// @Success      200 {string} string
// @Failure 	 200,204,400,401,404 {object} object
// @Router       /uom/deleteById/{id} [delete]
func (r *uomController) DeleteById(context *gin.Context) {
	id := context.Param("id")
	if strings.Trim(id, " ") == "" {
		context.AbortWithStatusJSON(http.StatusBadRequest, commons.ErrorResponse(constants.NoIdFound))
		return
	}

	if err := r.uomService.DeleteById(context, id); err != nil {
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
