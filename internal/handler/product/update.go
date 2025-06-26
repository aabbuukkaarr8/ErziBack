package product

import (
	"erzi_new/internal/service/product"
	"erzi_new/pkg/validator"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UpdateProduct struct {
	Title       *string  `json:"title"`
	Description *string  `json:"description"`
	Price       *float64 `json:"price"`
	Quantity    *int     `json:"quantity"`
	ImageURL    *string  `json:"image_url"`
}

func (m *UpdateProduct) ToSrv(id int) product.UpdateProduct {
	return product.UpdateProduct{
		ID:          id,
		Title:       m.Title,
		Description: m.Description,
		Price:       m.Price,
		Quantity:    m.Quantity,
	}
}

func (h *Handler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "неверный ID"})
		return
	}

	var input UpdateProduct
	err = validator.BindJSON(&input, c.Request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedSrv, err := h.srv.Update(input.ToSrv(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	result := Product{}
	result.FillFromService(updatedSrv)

	c.JSON(http.StatusOK, result)
}
