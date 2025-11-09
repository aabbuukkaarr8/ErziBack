package product

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouterDelete(msvc *MockService) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	h := NewHandler(msvc)
	r.DELETE("/products/:id", h.Delete)
	return r
}

func TestDelete_Success(t *testing.T) {
	msvc := new(MockService)
	router := setupRouterDelete(msvc)

	msvc.On("Delete", 7).Return(nil)

	req := httptest.NewRequest(http.MethodDelete, "/products/7", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
	assert.Empty(t, w.Body.String())

	msvc.AssertExpectations(t)
}

func TestDelete_ServiceError(t *testing.T) {
	msvc := new(MockService)
	router := setupRouterDelete(msvc)

	msvc.On("Delete", 7).Return(errors.New("delete failed"))

	req := httptest.NewRequest(http.MethodDelete, "/products/7", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)

	var body map[string]string
	_ = json.Unmarshal(w.Body.Bytes(), &body)
	assert.Equal(t, "delete failed", body["error"])

	msvc.AssertExpectations(t)
}
