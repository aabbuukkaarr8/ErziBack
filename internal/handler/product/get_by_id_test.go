package product

import (
	"encoding/json"
	"errors"
	svc "erzi_new/internal/service/product"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func setupRouterGetByID(msvc *MockService) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	h := NewHandler(msvc)
	r.GET("/products/:id", h.GetByID)
	return r
}

func TestGetByID_Success(t *testing.T) {
	msvc := new(MockService)
	router := setupRouterGetByID(msvc)

	now := time.Now().Truncate(time.Second)
	serviceProd := &svc.Model{
		ID:          7,
		Title:       "Tasty Water",
		Description: "Pure spring",
		Price:       2.50,
		ImageURL:    "http://img",
		Quantity:    100,
		CreatedAt:   now,
	}
	msvc.On("GetByID", 7).Return(serviceProd, nil)

	req := httptest.NewRequest(http.MethodGet, "/products/7", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp Model
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)

	var want Model
	want.FillFromService(serviceProd)
	assert.Equal(t, want, resp)

	msvc.AssertExpectations(t)
}

func TestGetByID_NotFound(t *testing.T) {
	msvc := new(MockService)
	router := setupRouterGetByID(msvc)

	msvc.On("GetByID", 7).Return((*svc.Model)(nil), errors.New("[Get By ID]Product Not Found"))

	req := httptest.NewRequest(http.MethodGet, "/products/7", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)

	var body map[string]string
	_ = json.Unmarshal(w.Body.Bytes(), &body)
	assert.Equal(t, "[Get By ID]Product Not Found", body["error"])

	msvc.AssertExpectations(t)
}
