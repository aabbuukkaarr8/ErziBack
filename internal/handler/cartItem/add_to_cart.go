package cartItem

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func (h *Handler) AddCartItem(c *gin.Context) {
	idStr := c.Param("product_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logrus.WithError(err).Errorf("[strconv.Atoi] Your id is not int")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	raw, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	s, ok := raw.(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	userID, err := uuid.Parse(s)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	i := AddCartItemRequest{
		ProductID: id,
		UserID:    userID,
	}
	iSrv := i.ToSrv()
	_, err = h.srv.Add(iSrv)
	if err != nil {
		logrus.WithError(err).Errorf("[Srv.Add] cant add item")
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, nil)
}
