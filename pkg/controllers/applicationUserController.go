package controllers

import (
	"errors"
	"net/http"
	"opslaundry/pkg/commons"
	"opslaundry/pkg/dto"
	"opslaundry/pkg/models"
	"opslaundry/pkg/services"
	"opslaundry/pkg/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ApplicationUserController interface {
	Register(c *gin.Context)
	Update(c *gin.Context)
	GetAll(context *gin.Context)
	GetProfile(c *gin.Context)
	// Register(context *gin.Context)
	// Lookup(context *gin.Context)
	// GetUserFirebase(context *gin.Context)
	// // UpdateProfile(context *gin.Context)
	// // ChangePhone(context *gin.Context)
	// ChangePassword(context *gin.Context)
	// // Profile(context *gin.Context)
	// GetPagination(context *gin.Context)
	// GetAll(context *gin.Context)
	// GetById(context *gin.Context)
	// GetViewById(context *gin.Context)
	// DeleteById(context *gin.Context)
	// RecoverPassword(context *gin.Context)

}

type applicationUserController struct {
	applicationUserService services.ApplicationUserService
	logger                 commons.OpsLogger
}

func NewApplicationUserController(db *gorm.DB) ApplicationUserController {
	return &applicationUserController{
		applicationUserService: services.NewApplicationUserService(db),
		logger:                 commons.NewLogger(),
	}
}

// @Tags         Application User
// @Security 	 BearerAuth
// @Summary
// @Description
// @Accept       json
// @Produce      json
// @Param        Request body dto.RegisterDto true "application_user"
// @Success      200 {object} object
// @Failure 	 204,400,401,404 {string} string
// @Router       /application_user/register [post]
func (r applicationUserController) Register(context *gin.Context) {
	var dto dto.RegisterDto

	if err := context.BindJSON(&dto); err != nil {
		context.JSON(http.StatusBadRequest, utils.ErrorToArray(err.Error()))
		return
	}

	result, err := r.applicationUserService.Register(context, dto)
	if err != nil {
		context.JSON(http.StatusBadRequest, commons.ErrorResponse(err.Error()))
		return
	}

	context.JSON(http.StatusOK, result)
}

// @Tags         Application User
// @Security 	 BearerAuth
// @Summary
// @Description
// @Accept       json
// @Produce      json
// @Param        Request body models.ApplicationUser true "application_user"
// @Success      200 {object} object
// @Failure 	 204,400,401,404 {string} string
// @Router       /application_user/update [patch]
func (r applicationUserController) Update(context *gin.Context) {
	//var dto RegisterDto
	var record models.ApplicationUser

	if err := context.BindJSON(&record); err != nil {
		context.JSON(http.StatusBadRequest, utils.ErrorToArray(err.Error()))
		return
	}

	result, err := r.applicationUserService.Update(context, record)
	if err != nil {
		context.JSON(http.StatusBadRequest, commons.ErrorResponse(err.Error()))
		return
	}

	context.JSON(http.StatusOK, result)
}

// @Tags         Application User
// @Security 	 BearerAuth
// @Summary
// @Description
// @Accept       json
// @Produce      json
// @Success      200 {object} object
// @Failure 	 204,400,401,404 {string} string
// @Router       /application_user/getProfile [get]
func (r applicationUserController) GetProfile(context *gin.Context) {
	result, err := r.applicationUserService.GetProfile(context)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.AbortWithStatus(http.StatusNoContent)
			return
		}
		context.JSON(http.StatusBadRequest, commons.ErrorResponse(err.Error()))
		return
	}

	context.JSON(http.StatusOK, result)
}

// @Tags         Application User
// @Security 	 BearerAuth
// @Summary
// @Description
// @Accept       json
// @Produce      json
// @Success      200 {object} object
// @Failure 	 204,400,401,404 {string} string
// @Router       /application_user/getAll [get]
func (r applicationUserController) GetAll(context *gin.Context) {
	result, err := r.applicationUserService.GetAll(context)
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

// // @Tags         Application User
// // @Security 	 BearerAuth
// // @Summary
// // @Description
// // @Accept       json
// // @Produce      json
// // @Param        Request body dto.ApplicationUserRegisterDto true " "
// // @Success      200 {object} helper.StandartResult
// // @Failure 	 400,404 {object} object
// // @Router       /application_user/register [post]
// func (r applicationUserController) Register(context *gin.Context) {
// 	r.logger.Info("Registering....")
// 	var recordDto dto.ApplicationUserRegisterDto
// 	errDto := context.Bind(&recordDto)
// 	if errDto != nil {
// 		res := helper.StandartResult{Success: false, Message: errDto.Error(), Data: nil}
// 		r.logger.Error(errDto.Error())
// 		context.JSON(http.StatusBadRequest, res)
// 		return
// 	}

// 	// authHeader := context.GetHeader("Authorization")
// 	// if authHeader == "" {
// 	// 	context.JSON(http.StatusUnauthorized, "Unauthorized")
// 	// 	return
// 	// }
// 	// _ = r.jwtService.GetUserByToken(authHeader)

// 	if recordDto.UserType == 1 && (recordDto.Email == "" || recordDto.Password == "") {
// 		res := helper.StandartResult{Success: false, Message: "Email dan Password tidak boleh kosong.", Data: nil}
// 		context.JSON(http.StatusBadRequest, res)
// 		return
// 	}

// 	// recordDto.CreatedBy = userIdentity.UserId
// 	// recordDto.OrganizationId = userIdentity.OrganizationId
// 	result := r.applicationUserService.Register(context, recordDto)
// 	if result.Success {
// 		result.Message = helper.RegisterUserSucceededMessage
// 	}

// 	// response := helper.StandartResult{result.Status, result.Message, result.Data}
// 	// response := helper.StandartResult{Success: false, Message: "Ok!", Data: nil}
// 	context.JSON(http.StatusOK, result)
// 	return
// }

// // @Tags         Application User
// // @Security 	 BearerAuth
// // @Summary
// // @Description
// // @Accept       json
// // @Produce      json
// // @Param        body body helper.ReactSelectRequest true " "
// // @Success      200 {object} helper.StandartResult
// // @Failure 	 400,404 {object} object
// // @Router       /application_user/lookup [post]
// func (r applicationUserController) Lookup(context *gin.Context) {
// 	var request helper.ReactSelectRequest

// 	errDto := context.Bind(&request)
// 	if errDto != nil {
// 		res := helper.StandartResult{Success: false, Message: errDto.Error(), Data: nil}
// 		context.JSON(http.StatusBadRequest, res)
// 		return
// 	}

// 	var result = r.applicationUserService.Lookup(request)
// 	context.JSON(http.StatusOK, result)
// 	return
// }

// // func (r applicationUserController) Lookup(context *gin.Context) {
// // 	var request helper.ReactSelectRequest
// // 	qry := context.Request.URL.Query()

// // 	if _, found := qry["q"]; found {
// // 		request.Q = fmt.Sprint(qry["q"][0])
// // 	}
// // 	request.Field = helper.StringifyToArray(fmt.Sprint(qry["field"]))

// // 	var result = c.applicationUserService.Lookup(request)
// // 	response := helper.StandartResult(true, "Ok", result.Data)
// // 	context.JSON(http.StatusOK, response)
// // }

// // @Tags         Application User
// // @Security 	 BearerAuth
// // @Accept       json
// // @Produce      json
// // @Param        body body commons.DataTableRequest true " "
// // @Success      200 {object} helper.StandartResult
// // @Failure 	 400,404 {object} object
// // @Router       /application_user/getPagination [post]
// func (r applicationUserController) GetPagination(ctx *gin.Context) {
// 	var req commons.DataTableRequest
// 	errDto := ctx.Bind(&req)
// 	if errDto != nil {
// 		res := helper.StandartResult{Success: false, Message: errDto.Error(), Data: nil}
// 		ctx.JSON(http.StatusBadRequest, res)
// 		return
// 	}

// 	var result = r.applicationUserService.GetPagination(ctx, req)
// 	ctx.JSON(http.StatusOK, result)
// 	return
// }

// // func (r applicationUserController) UpdateProfile(context *gin.Context) {
// // 	var dto dto.ApplicationUserDescriptionDto
// // 	errDTO := context.ShouldBind(&dto)

// // 	if errDTO != nil {
// // 		res := helper.StandartResult(false, errDTO.Error(), helper.EmptyObj{})
// // 		context.AbortWithStatusJSON(http.StatusBadRequest, res)
// // 	}

// // 	authHeader := context.GetHeader("Authorization")
// // 	token, errToken := c.jwtService.ValidateToken(authHeader)
// // 	if errToken != nil {
// // 		panic(errToken.Error())
// // 	}
// // 	claims := token.Claims.(jwt.MapClaims)

// // 	dto.UpdatedBy = fmt.Sprintf("%v", claims["user_id"])
// // 	result := c.applicationUserService.UpdateProfile(dto)
// // 	if result.Status {
// // 		response := helper.StandartResult(result.Status, "Ok", helper.EmptyObj{})
// // 		context.JSON(http.StatusOK, response)
// // 	} else {
// // 		response := helper.StandartResult(result.Status, result.Message, helper.EmptyObj{})
// // 		context.JSON(http.StatusOK, response)
// // 	}
// // }

// // func (r applicationUserController) Profile(context *gin.Context) {
// // 	authHeader := context.GetHeader("Authorization")
// // 	token, errToken := c.jwtService.ValidateToken(authHeader)
// // 	if errToken != nil {
// // 		panic(errToken.Error())
// // 	}

// // 	claims := token.Claims.(jwt.MapClaims)
// // 	id := fmt.Sprintf("%v", claims["user_id"])
// // 	user := c.applicationUserService.UserProfile(id)

// // 	res := helper.StandartResult(true, "Ok!", user)
// // 	context.JSON(http.StatusOK, res)
// // }

// // @Tags         Application User
// // @Security 	 BearerAuth
// // @Summary
// // @Description
// // @Accept       json
// // @Produce      json
// // @Param        id path string true "id"
// // @Success      200 {object} helper.StandartResult
// // @Failure 	 400,404 {object} object
// // @Router       /application_user/deleteById/{id} [delete]
// func (r applicationUserController) DeleteById(context *gin.Context) {
// 	id := context.Param("id")
// 	if id == "" {
// 		response := helper.StandartResult{Success: false, Message: helper.NoIdFound, Data: nil}
// 		context.JSON(http.StatusBadRequest, response)
// 		return
// 	}

// 	response := r.applicationUserService.DeleteById(id)
// 	context.JSON(http.StatusOK, response)
// 	return
// }

// // @Tags         Application User
// // @Security 	 BearerAuth
// // @Summary
// // @Description
// // @Accept       json
// // @Produce      json
// // @Param        id path string true "id"
// // @Success      200 {object} helper.StandartResult
// // @Failure 	 400,404 {object} object
// // @Router       /application_user/getById/{id} [get]
// func (r applicationUserController) GetById(context *gin.Context) {
// 	id := context.Param("id")
// 	if id == "" {
// 		response := helper.StandartResult{Success: false, Message: helper.NoIdFound, Data: nil}
// 		context.JSON(http.StatusBadRequest, response)
// 		return
// 	}

// 	response := r.applicationUserService.GetById(id)
// 	context.JSON(http.StatusOK, response)
// 	return
// }

// // @Tags         Application User
// // @Security 	 BearerAuth
// // @Summary
// // @Description
// // @Accept       json
// // @Produce      json
// // @Param        id path string true "id"
// // @Success      200 {object} helper.StandartResult
// // @Failure 	 400,404 {object} object
// // @Router       /application_user/getViewById/{id} [get]
// func (r applicationUserController) GetViewById(context *gin.Context) {
// 	id := context.Param("id")
// 	if id == "" {
// 		response := helper.StandartResult{Success: false, Message: helper.NoIdFound, Data: nil}
// 		context.JSON(http.StatusBadRequest, response)
// 		return
// 	}

// 	response := r.applicationUserService.GetViewById(id)
// 	context.JSON(http.StatusOK, response)
// 	return
// }

// // @Tags         Application User
// // @Security 	 BearerAuth
// // @Summary
// // @Description
// // @Accept       json
// // @Produce      json
// // @Param        Request body dto.ChangePasswordDto true " "
// // @Success      200 {object} helper.StandartResult
// // @Failure 	 400,404 {object} object
// // @Router       /application_user/changePassword [post]
// func (r applicationUserController) ChangePassword(context *gin.Context) {
// 	var recordDto dto.ChangePasswordDto
// 	errDto := context.Bind(&recordDto)
// 	if errDto != nil {
// 		res := helper.StandartResult{Success: false, Message: errDto.Error(), Data: nil}
// 		r.logger.Error(errDto.Error())
// 		context.JSON(http.StatusBadRequest, res)
// 		return
// 	}

// 	response := r.applicationUserService.ChangePassword(recordDto)
// 	context.JSON(http.StatusOK, response)
// 	return
// }

// // func (r applicationUserController) ChangePhone(context *gin.Context) {
// // 	var recordDto dto.ChangePhoneDto
// // 	err := context.ShouldBind(&recordDto)
// // 	if err != nil {
// // 		response := helper.StandartResult(false, err.Error(), helper.EmptyObj{})
// // 		context.JSON(http.StatusBadRequest, response)
// // 		return
// // 	}

// // 	result := c.applicationUserService.ChangePhone(recordDto)
// // 	context.JSON(http.StatusOK, result)
// // }

// // func (r applicationUserController) RecoverPassword(context *gin.Context) {
// // 	var recordDto dto.RecoverPasswordDto
// // 	err := context.ShouldBind(&recordDto)
// // 	if err != nil {
// // 		response := helper.StandartResult(false, err.Error(), helper.EmptyObj{})
// // 		context.JSON(http.StatusBadRequest, response)
// // 		return
// // 	}
// // 	result := c.applicationUserService.RecoverPassword(recordDto.Id, recordDto.OldPassword)
// // 	response := helper.StandartResult(true, "Ok!", result)
// // 	context.JSON(http.StatusOK, response)
// // }
