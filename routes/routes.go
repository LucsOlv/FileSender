package api

import (
	"filesender/docs"
	"filesender/handlers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Router gerencia todas as rotas da aplicação
type Router struct {
	controllers *handlers.Controllers
}

// NewRouter cria um novo router
func NewRouter(controllers *handlers.Controllers) *Router {
	return &Router{
		controllers: controllers,
	}
}

// SetupRoutes configura todas as rotas da aplicação
func (r *Router) SetupRoutes(engine *gin.Engine) {
	api := engine.Group("/api")
	// Configuração do Swagger
	docs.SwaggerInfo.BasePath = "/api"
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Rotas de arquivos
	api.POST("/upload", r.controllers.File.HandleUpload)
}
