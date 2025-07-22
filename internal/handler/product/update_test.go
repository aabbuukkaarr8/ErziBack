package product

import (
	"bytes"
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

func setupRouterUpdate(msvc *MockService) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	h := NewHandler(msvc)
	r.PUT("/products/:id", h.Update)
	return r
}

func ptrFloat64(f float64) *float64 { return &f }

func TestUpdate_Success(t *testing.T) {
	msvc := new(MockService)
	router := setupRouterUpdate(msvc)

	now := time.Now().Truncate(time.Second)
	returned := &svc.Model{
		ID:          7,
		Title:       "",
		Description: "",
		Price:       4.4,
		ImageURL:    "",
		Quantity:    0,
		Category:    "beverages",
		CreatedAt:   now,
	}
	dto := svc.UpdateProduct{ID: 7, Price: ptrFloat64(4.4)}
	msvc.On("Update", dto).Return(returned, nil)

	body, _ := json.Marshal(map[string]float64{"price": 4.4})
	req := httptest.NewRequest(http.MethodPut, "/products/7", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp Model
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))

	want := Model{}
	want.FillFromService(returned)
	assert.Equal(t, want, resp)

	msvc.AssertExpectations(t)
}

func TestUpdate_BadRequest(t *testing.T) {
	msvc := new(MockService)
	router := setupRouterUpdate(msvc)

	req := httptest.NewRequest(http.MethodPut, "/products/7", bytes.NewReader([]byte(`{price:bad}`)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	msvc.AssertNotCalled(t, "Update")
}

func TestUpdate_ServiceError(t *testing.T) {
	msvc := new(MockService)
	router := setupRouterUpdate(msvc)

	dto := svc.UpdateProduct{ID: 7, Price: ptrFloat64(4.4)}
	msvc.On("Update", dto).Return((*svc.Model)(nil), errors.New("update failed"))

	body, _ := json.Marshal(map[string]float64{"price": 4.4})
	req := httptest.NewRequest(http.MethodPut, "/products/7", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)

	var bodyMap map[string]string
	_ = json.Unmarshal(w.Body.Bytes(), &bodyMap)
	assert.Equal(t, "update failed", bodyMap["error"])

	msvc.AssertExpectations(t)
}
