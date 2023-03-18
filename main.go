package main

import (
	"context"
	"opslaundry/docs"
	"opslaundry/pkg/config/db"
	"opslaundry/routes"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	r := gin.Default()
	db, port := db.Connect()

	// programatically set swagger info
	docs.SwaggerInfo.Title = "Ops Laundry API Documentation"
	docs.SwaggerInfo.Description = "This is a full API documentation of Ops Laundry."
	docs.SwaggerInfo.Version = "1.0"
	// docs.SwaggerInfo.Host = *SERVER_NAME + ":" + *SERVER_PORT
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	//r.Use(GinContextToContextMiddleware())
	routes.RegisterRoutes(r, db)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(port)
}

func GinContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), "GinContextKey", c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
