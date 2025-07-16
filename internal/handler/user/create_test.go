package user

//
//import (
//	"bytes"
//	"encoding/json"
//	svc "erzi_new/internal/service/user"
//	"github.com/gin-gonic/gin"
//	"github.com/stretchr/testify/assert"
//	"github.com/stretchr/testify/mock"
//	"net/http"
//	"net/http/httptest"
//	"testing"
//)
//
//func SetupRouter(msvc *MockService) *gin.Engine {
//	gin.SetMode(gin.TestMode)
//	r := gin.New()
//	h := NewHandler(msvc)
//	r.POST("/user/create", h.Create)
//	return r
//}
//
//func TestHandler_Create(t *testing.T) {
//	msvc := new(MockService)
//	router := SetupRouter(msvc)
//	input := svc.CreateUser{
//		Username: "test",
//		Email:    "test@gmail.com",
//		Password: "test123",
//		Role:     "user",
//	}
//	body, _ := json.Marshal(input)
//	returned := &svc.User{
//		ID:       123,
//		Username: input.Username,
//		Email:    input.Email,
//		Password: input.Password,
//		Role:     input.Role,
//	}
//	msvc.On("Create", mock.MatchedBy(func(in svc.CreateUser) bool {
//		return in == input
//	})).Return(returned, nil)
//
//	req := httptest.NewRequest(http.MethodPost, "/user/create", bytes.NewReader(body))
//	req.Header.Set("Content-Type", "application/json")
//	w := httptest.NewRecorder()
//	router.ServeHTTP(w, req)
//
//	assert.Equal(t, http.StatusCreated, w.Code)
//	var resp svc.User
//	_ = json.Unmarshal(w.Body.Bytes(), &resp)
//	assert.Equal(t, *returned, resp)
//	msvc.AssertExpectations(t)
//
//}
