package routes

import (
	"opslaundry/pkg/controllers"
	"opslaundry/pkg/middleware"
	"opslaundry/pkg/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UomRoutes(r *gin.RouterGroup, db *gorm.DB, jwtService services.JWTService) {
	controller := controllers.NewUomController(db)
	routes := r.Group("/uom", middleware.AuthorizeJWT(jwtService))
	{
		routes.POST("/create", controller.Create)
		routes.PATCH("/update", controller.Update)
		routes.POST("/getPagination", controller.GetPagination)
		routes.GET("/getAll", controller.GetAll)
		routes.GET("/getById/:id", controller.GetById)
		routes.DELETE("/deleteById/:id", controller.DeleteById)
	}
}
