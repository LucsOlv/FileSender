package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

// Config armazena todas as configurações da aplicação
type Config struct {
	ServerPort       string
	MongoURI         string
	MongoDB          string
	MongoCollection  string
	RabbitMQURI      string
	RabbitExchange   string
	RabbitQueue      string
	RabbitRoutingKey string
	UploadDir        string
}

// LoadConfig carrega as configurações do arquivo .env e variáveis de ambiente
func LoadConfig() *Config {
	// Carregar variáveis do arquivo .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Arquivo .env não encontrado, usando variáveis de ambiente")
	}

	return &Config{
		ServerPort:       getEnv("SERVER_PORT", "8080"),
		MongoURI:         getEnv("MONGODB_URI", "mongodb://localhost:27017"),
		MongoDB:          getEnv("MONGODB_DATABASE", "filesender"),
		MongoCollection:  getEnv("MONGODB_COLLECTION", "files"),
		RabbitMQURI:      getEnv("RABBITMQ_URI", "amqp://guest:guest@localhost:5672/"),
		RabbitExchange:   getEnv("RABBITMQ_EXCHANGE", "filesender"),
		RabbitQueue:      getEnv("RABBITMQ_QUEUE", "file_uploads"),
		RabbitRoutingKey: getEnv("RABBITMQ_ROUTING_KEY", "file.uploaded"),
		UploadDir:        getEnv("UPLOAD_DIR", "./uploads"),
	}
}

// Função auxiliar para obter variáveis de ambiente
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// GetInt obtém uma variável de ambiente como inteiro
func GetInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if valueStr == "" {
		return defaultValue
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return defaultValue
	}

	return value
}

// GetDuration obtém uma variável de ambiente como duração em segundos
func GetDuration(key string, defaultValue time.Duration) time.Duration {
	valueStr := getEnv(key, "")
	if valueStr == "" {
		return defaultValue
	}

	valueInt, err := strconv.Atoi(valueStr)
	if err != nil {
		return defaultValue
	}

	return time.Duration(valueInt) * time.Second
}
