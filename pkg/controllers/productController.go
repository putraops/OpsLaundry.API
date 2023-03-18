package controllers

import (
	"errors"
	"net/http"
	"opslaundry/pkg/commons"
	"opslaundry/pkg/constants"
	"opslaundry/pkg/dto"
	"opslaundry/pkg/models"
	"opslaundry/pkg/services"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mashingan/smapping"
	"gorm.io/gorm"
)

type ProductController interface {
	Create(c *gin.Context)
	AddDetail(context *gin.Context)
	Update(c *gin.Context)
	GetPagination(context *gin.Context)
	GetAll(context *gin.Context)
	GetById(context *gin.Context)
	DeleteById(context *gin.Context)
}

type productController struct {
	productService       services.ProductService
	productDetailService services.ProductDetailService
	logger               commons.OpsLogger
}

func NewProductController(db *gorm.DB) ProductController {
	return &productController{
		productService:       services.NewProductService(db),
		productDetailService: services.NewProductDetailService(db),
		logger:               commons.NewLogger(),
	}
}

// @Tags         Product
// @Security 	 BearerAuth
// @Summary
// @Description
// @Accept       json
// @Produce      json
// @Param        Request body dto.NewProductDto true "Product"
// @Failure 	 200,204,400,401,404 {object} object
// @Router       /product/create [post]
func (r *productController) Create(context *gin.Context) {
	var dto dto.NewProductDto
	var record models.Product

	if err := context.BindJSON(&dto); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, commons.ErrorResponse(err.Error()))
		return
	}

	if err := smapping.FillStruct(&record, smapping.MapFields(&dto)); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	result, err := r.productService.Create(context, record)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, commons.ErrorResponse(err.Error()))
		return
	}

	context.JSON(http.StatusOK, result)
	return
}

// @Tags         Product
// @Security 	 BearerAuth
// @Summary
// @Description
// @Accept       json
// @Produce      json
// @Param        Request body models.ProductDetail true "Product Detail"
// @Failure 	 200,204,400,401,404 {object} object
// @Router       /product/detail/add [post]
func (r *productController) AddDetail(context *gin.Context) {
	var record models.ProductDetail

	if err := context.BindJSON(&record); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, commons.ErrorResponse(err.Error()))
		return
	}

	result, err := r.productDetailService.Create(context, record)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, commons.ErrorResponse(err.Error()))
		return
	}

	context.JSON(http.StatusOK, result)
	return
}

// @Tags         Product
// @Security 	 BearerAuth
// @Summary
// @Description
// @Accept       json
// @Produce      json
// @Param        Request body models.Product true "product"
// @Failure 	 200,204,400,401,404 {object} object
// @Router       /product/update [patch]
func (r *productController) Update(context *gin.Context) {
	var record models.Product

	if err := context.BindJSON(&record); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, commons.ErrorResponse(err.Error()))
		return
	}

	result, err := r.productService.Update(context, record)
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

// @Tags         Product
// @Security 	 BearerAuth
// @Accept       json
// @Produce      json
// @Param        body body commons.DataTableRequest true " "
// @Failure 	 200,204,400,401,404 {object} object
// @Router       /product/getPagination [post]
func (r *productController) GetPagination(context *gin.Context) {
	var req commons.DataTableRequest

	if err := context.BindJSON(&req); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, commons.ErrorResponse(err.Error()))
		return
	}

	result, err := r.productService.GetPagination(context, req)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, commons.ErrorResponse(err.Error()))
		return
	}
	context.JSON(http.StatusOK, result)
	return
}

// @Tags         Product
// @Security 	 BearerAuth
// @Summary
// @Description
// @Accept       json
// @Produce      json
// @Failure 	 200,204,400,401,404 {object} object
// @Router       /product/getAll [get]
func (r *productController) GetAll(context *gin.Context) {
	result, err := r.productService.GetAll(context)

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

// @Tags         Product
// @Security 	 BearerAuth
// @Summary
// @Description
// @Accept       json
// @Produce      json
// @Param        id path string true "id"
// @Failure 	 200,204,400,401,404 {object} object
// @Router       /product/getById/{id} [get]
func (r *productController) GetById(context *gin.Context) {
	id := context.Param("id")
	if strings.Trim(id, " ") == "" {
		context.AbortWithStatusJSON(http.StatusBadRequest, commons.ErrorResponse(constants.NoIdFound))
		return
	}

	result, err := r.productService.GetById(context, id)
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

// @Tags         Product
// @Security 	 BearerAuth
// @Summary
// @Description
// @Accept       json
// @Produce      json
// @Param        id path string true "id"
// @Success      200 {string} string
// @Failure 	 200,204,400,401,404 {object} object
// @Router       /product/deleteById/{id} [delete]
func (r *productController) DeleteById(context *gin.Context) {
	id := context.Param("id")
	if strings.Trim(id, " ") == "" {
		context.AbortWithStatusJSON(http.StatusBadRequest, commons.ErrorResponse(constants.NoIdFound))
		return
	}

	if err := r.productService.DeleteById(context, id); err != nil {
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
