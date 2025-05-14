package handlers

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestUploadFiles(t *testing.T) {
	// Configurar o modo de teste do Gin
	gin.SetMode(gin.TestMode)
	t.Run("Deve retornar erro quando array de arquivos está vazio", func(t *testing.T) {
		// Criar um request HTTP de teste com formulário vazio
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		// Simular uma requisição multipart com formulário mas sem arquivos
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		// Adicionar um campo qualquer para ter um formulário válido
		writer.WriteField("test", "value")
		writer.Close()

		req, _ := http.NewRequest("POST", "/api/form", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		c.Request = req

		// Configurar o contexto para ter um formulário multipart
		c.Request.MultipartForm = &multipart.Form{
			File: map[string][]*multipart.FileHeader{},
		}

		// Chamar o handler
		UploadFiles(c)

		// Verificar o resultado usando Testify
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "Nenhum arquivo foi enviado")
	})
}
