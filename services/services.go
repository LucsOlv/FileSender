package services

import (
	"filesender/config"
	rabbitmq "filesender/messaging"
)

// Services agrupa todos os serviços da aplicação
type Services struct {
	File *FileService
}

// NewServices cria e retorna todos os serviços da aplicação
func NewServices(rabbitMQ *rabbitmq.Publisher, config *config.Config) *Services {
	return &Services{
		File: NewFileService(rabbitMQ),
	}
}
