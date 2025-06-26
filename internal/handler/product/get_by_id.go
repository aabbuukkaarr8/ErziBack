package product

import (
	"net/http"
	"strconv"
	"time"

	"erzi_new/internal/service/product"
	"github.com/gin-gonic/gin"
)

type Product struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	ImageURL    string    `json:"image_url"`
	Quantity    int       `json:"quantity"`
	CreatedAt   time.Time `json:"created_at"`
}

func (m *Product) FillFromService(sm *product.Product) {
	m.ID = sm.ID
	m.Title = sm.Title
	m.Description = sm.Description
	m.Price = sm.Price
	m.ImageURL = sm.ImageURL
	m.Quantity = sm.Quantity
	m.CreatedAt = sm.CreatedAt
}

func (h *Handler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ps, err := h.srv.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	p := Product{}
	p.FillFromService(ps)
	c.JSON(http.StatusOK, p)
}
