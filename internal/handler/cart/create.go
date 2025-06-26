package cart

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) CreateCart(c *gin.Context) {
	var dto CreateCartDTO
	if err := c.BindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cartModel, err := h.srv.CreateCart(dto.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	resp := CartResponseDTO{ID: cartModel.ID, UserID: cartModel.UserID}
	c.JSON(http.StatusOK, resp)

}
