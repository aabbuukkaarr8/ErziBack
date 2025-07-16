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

func setupRouterGetAll(msvc *MockService) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	h := NewHandler(msvc)
	r.GET("/products", h.GetAll)
	return r
}

func TestGetAll_Success(t *testing.T) {
	msvc := new(MockService)
	router := setupRouterGetAll(msvc)

	now := time.Now().Truncate(time.Second)
	svcProducts := []svc.Product{
		{ID: 1, Title: "A", Description: "Desc A", Price: 1.1, ImageURL: "u1", Quantity: 10, Category: "cat1", CreatedAt: now},
		{ID: 2, Title: "B", Description: "Desc B", Price: 2.2, ImageURL: "u2", Quantity: 20, Category: "cat2", CreatedAt: now},
	}
	msvc.On("GetAll").Return(svcProducts, nil)

	req := httptest.NewRequest(http.MethodGet, "/products", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp []Product
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)

	want := make([]Product, len(svcProducts))
	for i, sp := range svcProducts {
		var p Product
		p.FillFromService(&sp)
		want[i] = p
	}
	assert.Equal(t, want, resp)

	msvc.AssertExpectations(t)
}

func TestGetAll_Error(t *testing.T) {
	msvc := new(MockService)
	router := setupRouterGetAll(msvc)

	msvc.On("GetAll").Return(nil, errors.New("db error"))

	req := httptest.NewRequest(http.MethodGet, "/products", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)

	var body map[string]string
	_ = json.Unmarshal(w.Body.Bytes(), &body)
	assert.Equal(t, "db error", body["error"])

	msvc.AssertExpectations(t)
}
