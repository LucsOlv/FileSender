package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"filesender/routes"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// Definir uma suite de testes
type APITestSuite struct {
	suite.Suite
	router *gin.Engine
}

// Configuração que roda antes de cada teste
func (s *APITestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
	s.router = gin.New()
	routes.SetupRoutes(s.router)
}

// Teste para o endpoint de upload
func (s *APITestSuite) TestUploadEndpoint() {
	// Criar uma requisição sem arquivos
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/form", nil)
	s.router.ServeHTTP(w, req)

	// Verificar resposta
	assert.Equal(s.T(), http.StatusBadRequest, w.Code)
}

// Teste para o Swagger UI
func (s *APITestSuite) TestSwaggerUI() {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/swagger", nil)
	s.router.ServeHTTP(w, req)

	// Verificar se a página do Swagger está acessível
	assert.NotEqual(s.T(), http.StatusNotFound, w.Code)
}

// Função para executar a suite de testes
func TestAPISuite(t *testing.T) {
	suite.Run(t, new(APITestSuite))
}

// Teste simples sem usar suite
func TestSimpleAPI(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	routes.SetupRoutes(router)

	t.Run("Endpoint de upload deve rejeitar requisições sem arquivos", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/form", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
