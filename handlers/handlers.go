// internal/api/handlers/handlers.go
package handlers

import (
	"filesender/config"
	rabbitmq "filesender/messaging"
	"filesender/services"
)

type Controllers struct {
	File *UploadHandler
}

// NewControllers cria e retorna todos os handlers da aplicação
func NewControllers(
	fileService *services.FileService,
	rabbitmq *rabbitmq.Publisher,
	config *config.Config) *Controllers {
	return &Controllers{
		File: NewUploadHandler(fileService, rabbitmq),
	}
}
