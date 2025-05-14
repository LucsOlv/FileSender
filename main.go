package main

import (
	"net/http"

	_ "filesender/docs" // Substitua "seu-pacote" pelo nome do seu módulo

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           API de Upload de Arquivos
// @version         1.0
// @description     API para upload de arquivos
// @host            localhost:8080
// @BasePath        /api

func main() {
	router := gin.Default()

	// Configuração do Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// --- Suas rotas da API ---
	apiV1 := router.Group("/api")
	{
		apiV1.POST("/form", uploadFiles)
	}

	router.Run(":8080")
}

// @Summary      Upload de arquivos
// @Description  Faz upload de múltiplos arquivos
// @Tags         uploads
// @Accept       multipart/form-data
// @Produce      json
// @Param        files  formData  file  true  "Arquivos para upload (múltiplos)"
// @Success      200    {object}  map[string]interface{}
// @Failure      400    {string}  string  "Nenhum arquivo foi enviado"
// @Router       /form [post]
func uploadFiles(c *gin.Context) {
	var form, err = c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, "Erro ao processar o formulário")
		return
	}

	files := form.File["files"]
	if len(files) == 0 {
		c.String(http.StatusBadRequest, "Nenhum arquivo foi enviado.")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Arquivos recebidos com sucesso"})
}
