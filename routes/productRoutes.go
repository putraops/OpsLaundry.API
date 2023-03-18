package routes

import (
	"opslaundry/pkg/controllers"
	"opslaundry/pkg/middleware"
	"opslaundry/pkg/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ProductRoutes(r *gin.RouterGroup, db *gorm.DB, jwtService services.JWTService) {
	controller := controllers.NewProductController(db)
	routes := r.Group("/product", middleware.AuthorizeJWT(jwtService))
	{
		routes.POST("/create", controller.Create)
		routes.POST("/detail/add", controller.AddDetail)
		routes.PATCH("/update", controller.Update)
		routes.POST("/getPagination", controller.GetPagination)
		routes.GET("/getAll", controller.GetAll)
		routes.GET("/getById/:id", controller.GetById)
		routes.DELETE("/deleteById/:id", controller.DeleteById)
	}
}
