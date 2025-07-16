package product

import (
	"github.com/sirupsen/logrus"
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
	Category    string    `json:"category"`
	CreatedAt   time.Time `json:"created_at"`
}

func (m *Product) FillFromService(sm *product.Product) {
	m.ID = sm.ID
	m.Title = sm.Title
	m.Description = sm.Description
	m.Price = sm.Price
	m.ImageURL = sm.ImageURL
	m.Quantity = sm.Quantity
	m.Category = sm.Category
	m.CreatedAt = sm.CreatedAt
}

func (h *Handler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logrus.WithError(err).Errorf("[strconv.Atoi] Your ID is not a number")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ps, err := h.srv.GetByID(id)
	if err != nil {
		logrus.WithError(err).Errorf("[Get By ID]Product Not Found")
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	p := Product{}
	p.FillFromService(ps)
	c.JSON(http.StatusOK, p)
}
