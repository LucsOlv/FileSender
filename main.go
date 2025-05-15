package main

import (
	"context"
	"filesender/handlers"
	api "filesender/routes"
	"filesender/services"

	"log"
	"net/http"

	"filesender/config"
	messaging "filesender/messaging"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		// Provide all components
		fx.Provide(
			config.LoadConfig,
			messaging.NewPublisher,

			// Serviços
			services.NewFileService,
			services.NewServices,

			// Controladores (todos os handlers)
			handlers.NewControllers,

			// Router
			api.NewRouter,

			// Gin engine
			func() *gin.Engine {
				return gin.Default()
			},
		),
		// Register lifecycle hooks
		fx.Invoke(registerHooks),
	)

	app.Run()
}

func registerHooks(
	lifecycle fx.Lifecycle,
	router *api.Router,
	engine *gin.Engine,
	config *config.Config,
	rabbitMQ *messaging.Publisher,
) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			// Setup routes
			router.SetupRoutes(engine)

			// Start HTTP server
			go func() {
				if err := engine.Run(":" + config.ServerPort); err != nil && err != http.ErrServerClosed {
					log.Fatalf("Failed to start server: %v", err)
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			// Close RabbitMQ connection
			return rabbitMQ.Close()
		},
	})
}
