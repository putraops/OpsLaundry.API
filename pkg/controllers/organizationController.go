package controllers

import (
	"net/http"
	"opslaundry/pkg/commons"
	"opslaundry/pkg/models"
	"opslaundry/pkg/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrganizationController interface {
	Update(c *gin.Context)
	// Lookup(context *gin.Context)
	// GetPagination(context *gin.Context)
	// GetAll(context *gin.Context)
	// GetById(context *gin.Context)
	// DeleteById(context *gin.Context)
}

type organizationController struct {
	organizationService services.OrganizationService
	logger              commons.OpsLogger
}

func NewOrganizationController(db *gorm.DB, jwtService services.JWTService) OrganizationController {
	return &organizationController{
		organizationService: services.NewOrganizationService(db, jwtService),
		logger:              commons.NewLogger(),
	}
}

// @Tags         Organization
// @Security 	 BearerAuth
// @Summary
// @Description
// @Accept       json
// @Produce      json
// @Param        Request body models.Organization true "organization"
// @Success      200 {object} object
// @Failure 	 204,400,401,404 {string} string
// @Router       /organization/update [patch]
func (r organizationController) Update(context *gin.Context) {
	var record models.Organization

	if err := context.BindJSON(&record); err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
		return
	}

	result, err := r.organizationService.Update(context, record)
	if err != nil {
		context.JSON(http.StatusBadRequest, commons.ErrorResponse(err.Error()))
		return
	}

	context.JSON(http.StatusOK, result)
}

// // @Tags         Organization
// // @Security 	 BearerAuth
// // @Summary
// // @Description
// // @Accept       json
// // @Produce      json
// // @Success      200 {object} helper.StandartResult
// // @Failure 	 400,404 {object} object
// // @Router       /organization/getAll [get]
// func (r organizationController) GetAll(context *gin.Context) {
// 	qry := context.Request.URL.Query()
// 	filter := make(map[string]interface{})

// 	for k, v := range qry {
// 		filter[k] = v
// 	}

// 	var response = r.organizationService.GetAll(context, filter)
// 	context.JSON(http.StatusOK, response)
// 	return
// }

// // @Tags         Organization
// // @Security 	 BearerAuth
// // @Summary
// // @Description
// // @Accept       json
// // @Produce      json
// // @Param        body body helper.ReactSelectRequest true " "
// // @Success      200 {object} helper.StandartResult
// // @Failure 	 400,404 {object} object
// // @Router       /organization/lookup [post]
// func (r organizationController) Lookup(context *gin.Context) {
// 	var request helper.ReactSelectRequest

// 	errDto := context.Bind(&request)
// 	if errDto != nil {
// 		res := helper.StandartResult{Success: false, Message: errDto.Error(), Data: nil}
// 		context.JSON(http.StatusBadRequest, res)
// 		return
// 	}

// 	var result = r.organizationService.Lookup(context, request)
// 	context.JSON(http.StatusOK, result)
// 	return
// }

// // @Tags         Organization
// // @Security 	 BearerAuth
// // @Accept       json
// // @Produce      json
// // @Param        body body commons.DataTableRequest true " "
// // @Success      200 {object} helper.StandartResult
// // @Failure 	 400,404 {object} object
// // @Router       /organization/getPagination [post]
// func (r organizationController) GetPagination(ctx *gin.Context) {
// 	var req commons.DataTableRequest
// 	errDto := ctx.Bind(&req)
// 	if errDto != nil {
// 		res := helper.StandartResult{Success: false, Message: errDto.Error(), Data: nil}
// 		ctx.JSON(http.StatusBadRequest, res)
// 		return
// 	}

// 	var result = r.organizationService.GetPagination(ctx, req)
// 	ctx.JSON(http.StatusOK, result)
// 	return
// }

// // @Tags         Organization
// // @Security 	 BearerAuth
// // @Summary
// // @Description
// // @Accept       json
// // @Produce      json
// // @Param        id path string true "id"
// // @Success      200 {object} helper.StandartResult
// // @Failure 	 400,404 {object} object
// // @Router       /organization/deleteById/{id} [delete]
// func (r organizationController) DeleteById(context *gin.Context) {
// 	id := context.Param("id")
// 	if id == "" {
// 		response := helper.StandartResult{Success: false, Message: helper.NoIdFound, Data: nil}
// 		context.JSON(http.StatusBadRequest, response)
// 		return
// 	}

// 	response := r.organizationService.DeleteById(context, id)
// 	context.JSON(http.StatusOK, response)
// 	return
// }

// // @Tags         Organization
// // @Security 	 BearerAuth
// // @Summary
// // @Description
// // @Accept       json
// // @Produce      json
// // @Param        id path string true "id"
// // @Success      200 {object} helper.StandartResult
// // @Failure 	 400,404 {object} object
// // @Router       /organization/getById/{id} [get]
// func (r organizationController) GetById(context *gin.Context) {
// 	id := context.Param("id")
// 	if id == "" {
// 		response := helper.StandartResult{Success: false, Message: helper.NoIdFound, Data: nil}
// 		context.JSON(http.StatusBadRequest, response)
// 		return
// 	}

// 	response := r.organizationService.GetById(context, id)
// 	context.JSON(http.StatusOK, response)
// 	return
// }
