package services

import (
	rabbitmq "filesender/messaging"
)

type FileService struct {
	rabbitMQ *rabbitmq.Publisher
}

func NewFileService(rabbitMQ *rabbitmq.Publisher) *FileService {
	return &FileService{
		rabbitMQ: rabbitMQ,
	}
}
