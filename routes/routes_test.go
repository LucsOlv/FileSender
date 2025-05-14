package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSetupRoutes(t *testing.T) {
	// Configurar o modo de teste do Gin
	gin.SetMode(gin.TestMode)

	// Criar um router de teste
	router := gin.New()
	SetupRoutes(router)

	t.Run("Deve ter rota para Swagger", func(t *testing.T) {
		// Criar um request HTTP de teste para a rota do Swagger
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/swagger", nil)
		router.ServeHTTP(w, req)

		// Verificar se a rota existe (não retorna 404)
		assert.NotEqual(t, http.StatusNotFound, w.Code)
	})

	t.Run("Deve ter rota para upload de arquivos", func(t *testing.T) {
		// Criar um request HTTP de teste para a rota de upload
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/form", nil)
		router.ServeHTTP(w, req)

		// Verificar se a rota existe (não retorna 404)
		assert.NotEqual(t, http.StatusNotFound, w.Code)
		// Deve retornar 400 porque não enviamos arquivos
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
