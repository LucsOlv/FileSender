package handlers

import (
	"io"
	"log"
	"net/http"

	rabbitmq "filesender/messaging"
	"filesender/services"

	"github.com/gin-gonic/gin"
)

type UploadHandler struct {
	fileService *services.FileService
	publisher   *rabbitmq.Publisher
}

func NewUploadHandler(fileService *services.FileService, publisher *rabbitmq.Publisher) *UploadHandler {
	return &UploadHandler{
		publisher: publisher,
	}
}

// @Summary      Upload de arquivos
// @Description  Faz upload de múltiplos arquivos
// @Tags         uploads
// @Accept       multipart/form-data
// @Produce      json
// @Param        files  formData  []file  true  "Arquivos para upload (múltiplos)"
// @Success      200    {object}  map[string]interface{} "Retorna os nomes dos arquivos enviados"
// @Failure      400    {object}  map[string]string
// @Failure      500    {object}  map[string]string
// @Router       /upload [post]
func (h *UploadHandler) HandleUpload(c *gin.Context) {
	// Adicione logs para debug
	log.Println("Recebendo requisição de upload")

	form, err := c.MultipartForm()
	if err != nil {
		log.Printf("Erro ao processar formulário: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao processar o formulário: " + err.Error()})
		return
	}

	files := form.File["files"]
	if len(files) == 0 {
		log.Println("Nenhum arquivo enviado")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nenhum arquivo foi enviado."})
		return
	}

	log.Printf("Recebidos %d arquivos", len(files))

	for i, file := range files {
		log.Printf("Processando arquivo %d: %s", i+1, file.Filename)

		fileHeader, err := file.Open()
		if err != nil {
			log.Printf("Erro ao abrir arquivo: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao abrir o arquivo: " + err.Error()})
			return
		}
		defer fileHeader.Close()

		// Ler o arquivo para o buffer
		fileBytes, err := io.ReadAll(fileHeader)
		if err != nil {
			log.Printf("Erro ao ler arquivo: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao ler o arquivo: " + err.Error()})
			return
		}

		log.Printf("Enviando arquivo %s para RabbitMQ", file.Filename)
		err = h.publisher.PublishFile(fileBytes, file.Filename)
		if err != nil {
			log.Printf("Erro ao publicar no RabbitMQ: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao enviar o arquivo para processamento: " + err.Error()})
			return
		}
		log.Printf("Arquivo %s enviado com sucesso", file.Filename)
	}

	log.Println("Todos os arquivos processados com sucesso")
	c.JSON(http.StatusOK, gin.H{"message": "Arquivos recebidos com sucesso"})
}
