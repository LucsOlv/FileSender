package routes

import (
	"filesender/docs"
	"filesender/handlers"

	rabbitmq "filesender/messaging"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(router *gin.Engine, publisher *rabbitmq.Publisher) {
	uploadHandler := handlers.NewUploadHandler(publisher)

	// Configuração do Swagger
	docs.SwaggerInfo.BasePath = "/api"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Endpoint de teste
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Rotas da API
	apiV1 := router.Group("/api")
	{
		// Endpoint de teste dentro do grupo API
		apiV1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "api pong",
			})
		})

		apiV1.POST("/form", uploadHandler.HandleUpload)
	}
}
