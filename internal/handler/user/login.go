package user

import (
	"erzi_new/pkg/validator"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) Login(c *gin.Context) {
	var req LoginRequest
	if err := validator.BindJSON(&req, c.Request); err != nil {
		logrus.WithError(err).Warn("Login: invalid request JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.srv.Login(req.Email, req.Password)
	if err != nil {
		logrus.WithError(err).Warn("Login: invalid credentials")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, LoginResponse{Token: token})
}
