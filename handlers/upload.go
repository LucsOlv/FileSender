package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary      Upload de arquivos
// @Description  Faz upload de múltiplos arquivos
// @Tags         uploads
// @Accept       multipart/form-data
// @Produce      json
// @Param        files  formData  file  true  "Arquivos para upload (múltiplos)"
// @Success      200    {object}  map[string]interface{}
// @Failure      400    {string}  string  "Nenhum arquivo foi enviado"
// @Router       /form [post]
func UploadFiles(c *gin.Context) {
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
