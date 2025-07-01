package product

import (
	"bytes"
	"encoding/json"
	svc "erzi_new/internal/service/product"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetupRouter(msvc *MockService) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	h := NewHandler(msvc)
	r.POST("/products/create", h.Create)
	return r
}

func TestCreate_Success(t *testing.T) {
	msvc := new(MockService)
	router := SetupRouter(msvc)

	input := svc.CreateProduct{
		Title:       "Water",
		Description: "Fresh spring water",
		Price:       1.99,
		Quantity:    10,
		Category:    "beverages",
	}
	body, _ := json.Marshal(input)

	returned := &svc.Product{
		ID:          123,
		Title:       input.Title,
		Description: input.Description,
		Price:       input.Price,
		Quantity:    input.Quantity,
		Category:    input.Category,
	}

	msvc.On("Create", input).Return(returned, nil)

	req := httptest.NewRequest(http.MethodPost, "/products/create", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	var resp svc.Product
	_ = json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, *returned, resp)

	msvc.AssertExpectations(t)
}

func TestCreate_ServiceError(t *testing.T) {
	msvc := new(MockService)
	router := SetupRouter(msvc)

	input := svc.CreateProduct{
		Title:       "Water",
		Description: "Spring",
		Price:       1.23,
		Quantity:    5,
		Category:    "beverages",
	}

	msvc.On("Create", input).Return((*svc.Product)(nil), assert.AnError)

	body, _ := json.Marshal(input)
	req := httptest.NewRequest(http.MethodPost, "/products/create", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	msvc.AssertExpectations(t)
}

func TestCreate_BadRequest(t *testing.T) {
	msvc := new(MockService)
	router := SetupRouter(msvc)

	req := httptest.NewRequest(http.MethodPost, "/products/create",
		bytes.NewReader([]byte(`{"title":"OnlyTitle"}`)),
	)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	msvc.AssertNotCalled(t, "Create", mock.Anything)
}
