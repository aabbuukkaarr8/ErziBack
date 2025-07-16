package user

import (
	userservice "erzi_new/internal/service/user"
	"erzi_new/pkg/validator"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (m *RegisterRequest) ToSrv() userservice.CreateUser {
	return userservice.CreateUser{
		Username: m.Username,
		Password: m.Password,
		Email:    m.Email,
		Role:     m.Role,
	}
}

func (m *User) FillFromService(srv *userservice.User) {
	m.ID = srv.ID
	m.Username = srv.Username
	m.Password = srv.Password
	m.Email = srv.Email
	m.Role = srv.Role
	m.CreatedAt = srv.CreatedAt
}

func (h *Handler) Create(c *gin.Context) {
	var req RegisterRequest
	err := validator.BindJSON(&req, c.Request)
	if err != nil {
		logrus.WithError(err).Warn("[Validator] Invalid JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": "неверные входные данные"})
		return
	}
	reqSrv := req.ToSrv()
	createdreqSrv, err := h.srv.Create(reqSrv)
	if err != nil {
		logrus.WithError(err).Warn("[Service] Failed to create service")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	createdReq := User{}
	createdReq.FillFromService(createdreqSrv)
	if createdReq.Role != "admin" {
		createdReq.Role = "user"
	}
	c.JSON(http.StatusCreated, createdReq)
}
