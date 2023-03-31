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

type ProductCategoryController interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	GetPagination(context *gin.Context)
	GetAll(context *gin.Context)
	GetById(context *gin.Context)
	DeleteById(context *gin.Context)
}

type productCategoryController struct {
	productCategoryService services.ProductCategoryService
	logger                 commons.OpsLogger
}

func NewProductCategoryController(db *gorm.DB) ProductCategoryController {
	return &productCategoryController{
		productCategoryService: services.NewProductCategoryService(db),
		logger:                 commons.NewLogger(),
	}
}

// @Tags         Product Category
// @Security 	 BearerAuth
// @Summary
// @Description
// @Accept       json
// @Produce      json
// @Param        Request body models.ProductCategory true "ProductCategory"
// @Failure 	 200,201,204,400,401,404 {object} object
// @Router       /product_category/create [post]
func (r *productCategoryController) Create(context *gin.Context) {
	var record models.ProductCategory

	if err := context.BindJSON(&record); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, commons.ErrorResponse(err.Error()))
		return
	}

	result, err := r.productCategoryService.Create(context, record)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, commons.ErrorResponse(err.Error()))
		return
	}

	context.JSON(http.StatusCreated, result)
}

// @Tags         Product Category
// @Security 	 BearerAuth
// @Summary
// @Description
// @Accept       json
// @Produce      json
// @Param        Request body models.ProductCategory true "productCategory"
// @Failure 	 200,204,400,401,404 {object} object
// @Router       /product_category/update [patch]
func (r *productCategoryController) Update(context *gin.Context) {
	var record models.ProductCategory

	if err := context.BindJSON(&record); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, commons.ErrorResponse(err.Error()))
		return
	}

	result, err := r.productCategoryService.Update(context, record)
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

// @Tags         Product Category
// @Security 	 BearerAuth
// @Accept       json
// @Produce      json
// @Param        body body commons.DataTableRequest true " "
// @Failure 	 200,204,400,401,404 {object} object
// @Router       /product_category/getPagination [post]
func (r *productCategoryController) GetPagination(context *gin.Context) {
	var req commons.DataTableRequest

	if err := context.BindJSON(&req); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	result := r.productCategoryService.GetPagination(context, req).(commons.DataTableResponse)
	context.JSON(http.StatusOK, result)
}

// @Tags         Product Category
// @Security 	 BearerAuth
// @Summary
// @Description
// @Accept       json
// @Produce      json
// @Failure 	 200,204,400,401,404 {object} object
// @Router       /product_category/getAll [get]
func (r *productCategoryController) GetAll(context *gin.Context) {
	result, err := r.productCategoryService.GetAll(context)

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

// @Tags         Product Category
// @Security 	 BearerAuth
// @Summary
// @Description
// @Accept       json
// @Produce      json
// @Param        id path string true "id"
// @Failure 	 200,204,400,401,404 {object} object
// @Router       /product_category/getById/{id} [get]
func (r *productCategoryController) GetById(context *gin.Context) {
	id := context.Param("id")
	if strings.Trim(id, " ") == "" {
		context.AbortWithStatusJSON(http.StatusBadRequest, commons.ErrorResponse(constants.NoIdFound))
		return
	}

	result, err := r.productCategoryService.GetById(context, id)
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

// @Tags         Product Category
// @Security 	 BearerAuth
// @Summary
// @Description
// @Accept       json
// @Produce      json
// @Param        id path string true "id"
// @Success      200 {string} string
// @Failure 	 200,204,400,401,404 {object} object
// @Router       /product_category/deleteById/{id} [delete]
func (r *productCategoryController) DeleteById(context *gin.Context) {
	id := context.Param("id")
	if strings.Trim(id, " ") == "" {
		context.AbortWithStatusJSON(http.StatusBadRequest, commons.ErrorResponse(constants.NoIdFound))
		return
	}

	if err := r.productCategoryService.DeleteById(context, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.AbortWithStatus(http.StatusNoContent)
			return
		}
		context.AbortWithStatusJSON(http.StatusBadRequest, commons.ErrorResponse(err.Error()))
		return
	}

	context.JSON(http.StatusOK, constants.DeletedMessage)
}
