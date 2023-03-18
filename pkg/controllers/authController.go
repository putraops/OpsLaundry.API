package controllers

import (
	"net/http"
	"opslaundry/pkg/commons"
	"opslaundry/pkg/dto"
	"opslaundry/pkg/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthController interface {
	Login(ctx *gin.Context)
}

type authController struct {
	db          *gorm.DB
	authService services.AuthService
	jwtService  services.JWTService
}

func NewAuthController(db *gorm.DB, jwtService services.JWTService) AuthController {
	return &authController{
		db:          db,
		jwtService:  jwtService,
		authService: services.NewAuthService(db, jwtService),
	}
}

// @Tags         Authentication
// @Security 	 ApiKeyAuth
// @Summary
// @Description
// @Accept       json
// @Produce      json
// @Param        Request body dto.LoginDto true "body"
// @Success      200 {object} object
// @Failure 	 400,401,404 {string} string
// @Router       /auth/login [post]
func (r authController) Login(c *gin.Context) {
	var loginDto dto.LoginDto
	err := c.ShouldBind(&loginDto)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	data, err := r.authService.VerifyCredential(loginDto)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, commons.ErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, data)
	return
}
