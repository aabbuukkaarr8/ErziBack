package cartItem

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) DeleteAll(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		logrus.Errorf("[c.Get] не найден userID в токене")
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	var intUserID int
	switch v := userID.(type) {
	case float64:
		intUserID = int(v)
	case int:
		intUserID = v
	default:
		c.JSON(401, gin.H{"[UserID]": "invalid user id type"})
		return
	}

	err := h.srv.DeleteAll(intUserID)
	if err != nil {
		logrus.WithError(err).Errorf("[srv.DeleteAll] Error")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}
