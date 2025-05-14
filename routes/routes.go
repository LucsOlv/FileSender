package routes

import (
	"filesender/docs"
	"filesender/handlers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetupRoutes configura todas as rotas da aplicação
func SetupRoutes(router *gin.Engine) {
	// Configuração do Swagger
	docs.SwaggerInfo.BasePath = "/api"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Rotas da API
	apiV1 := router.Group("/api")
	{
		apiV1.POST("/form", handlers.UploadFiles)
	}
}
