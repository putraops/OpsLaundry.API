package routes

import (
	"opslaundry/pkg/controllers"
	"opslaundry/pkg/middleware"
	"opslaundry/pkg/services"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

var (
	jwtService     services.JWTService
	authController controllers.AuthController
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	jwtService = services.NewJWTService()

	apiRoutes := r.Group("api")
	{
		authController = controllers.NewAuthController(db, jwtService)
		authRoutes := apiRoutes.Group("/auth")
		{
			authRoutes.POST("/login", authController.Login)
		}

		organizationController := controllers.NewOrganizationController(db, jwtService)
		organizationRoutes := apiRoutes.Group("/organization", middleware.AuthorizeJWT(jwtService))
		{
			organizationRoutes.PATCH("/update", organizationController.Update)
		}

		applicationUserController := controllers.NewApplicationUserController(db)
		applicationUserRoutes := apiRoutes.Group("/application_user", middleware.AuthorizeJWT(jwtService))
		{
			applicationUserRoutes.POST("/register", applicationUserController.Register)
			applicationUserRoutes.PATCH("/update", applicationUserController.Update)
			applicationUserRoutes.GET("/getAll", applicationUserController.GetAll)
			applicationUserRoutes.GET("/getProfile", applicationUserController.GetProfile)
		}

		tenantController := controllers.NewTenantController(db)
		tenantRoutes := apiRoutes.Group("/tenant", middleware.AuthorizeJWT(jwtService))
		{
			tenantRoutes.POST("/create", tenantController.Create)
			tenantRoutes.PATCH("/update", tenantController.Update)
			tenantRoutes.POST("/getPagination", tenantController.GetPagination)
			tenantRoutes.GET("/getAll", tenantController.GetAll)
			tenantRoutes.GET("/getById/:id", tenantController.GetById)
			tenantRoutes.DELETE("/deleteById/:id", tenantController.DeleteById)
		}

		productCategoryController := controllers.NewProductCategoryController(db)
		productCategoryRoutes := apiRoutes.Group("/product_category", middleware.AuthorizeJWT(jwtService))
		{
			productCategoryRoutes.POST("/create", productCategoryController.Create)
			productCategoryRoutes.PATCH("/update", productCategoryController.Update)
			productCategoryRoutes.POST("/getPagination", productCategoryController.GetPagination)
			productCategoryRoutes.GET("/getAll", productCategoryController.GetAll)
			productCategoryRoutes.GET("/getById/:id", productCategoryController.GetById)
			productCategoryRoutes.DELETE("/deleteById/:id", productCategoryController.DeleteById)
		}

		serviceTypeController := controllers.NewServiceTypeController(db)
		serviceTypeRoutes := apiRoutes.Group("/service_type", middleware.AuthorizeJWT(jwtService))
		{
			serviceTypeRoutes.POST("/create", serviceTypeController.Create)
			serviceTypeRoutes.PATCH("/update", serviceTypeController.Update)
			serviceTypeRoutes.POST("/getPagination", serviceTypeController.GetPagination)
			serviceTypeRoutes.GET("/getAll", serviceTypeController.GetAll)
			serviceTypeRoutes.GET("/getById/:id", serviceTypeController.GetById)
			serviceTypeRoutes.DELETE("/deleteById/:id", serviceTypeController.DeleteById)
		}

		UomRoutes(apiRoutes, db, jwtService)
		ProductRoutes(apiRoutes, db, jwtService)
	}
}
