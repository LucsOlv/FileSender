package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"filesender/config"
	_ "filesender/docs" // Substitua "seu-pacote" pelo nome do seu módulo
	"filesender/routes"

	rabbitmq "filesender/messaging"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// @title           API de Upload de Arquivos
// @version         1.0
// @description     API para upload de arquivos
// @host            localhost:8080
// @BasePath        /api

func main() {
	cfg := config.LoadConfig()

	if cfg == nil {
		log.Fatalf("Failed to load configuration")
	}
	router := gin.Default()
	gin.SetMode(gin.DebugMode)

	publisher, err := rabbitmq.NewPublisher(cfg.RabbitMQURI, "file_queue")
	if err != nil {
		log.Fatalf("Failed to create RabbitMQ publisher: %v", err)
	}

	// Configurar CORS
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	router.Use(cors.New(config))

	routes.SetupRoutes(router, publisher)

	srv := &http.Server{
		Addr:         ":" + cfg.ServerPort,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Iniciar servidor em uma goroutine
	go func() {
		log.Printf("Servidor iniciado na porta %s", cfg.ServerPort)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Erro ao iniciar servidor: %v", err)
		}
	}()

	// Configurar graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Desligando servidor...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Erro ao desligar servidor: %v", err)
	}

	log.Println("Servidor encerrado com sucesso")
}
