package product

import (
	"erzi_new/internal/service/product"
	"erzi_new/pkg/validator"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (a *CreateAttributeRequest) ToSrv() product.AttributeInput {
	return product.AttributeInput{
		ProductID: a.ProductID,
		Key:       a.Key,
		Value:     a.Value,
	}
}

func (m *Attribute) FillFromSRV(sm *product.Attribute) {
	m.Key = sm.Key
	m.Value = sm.Value
}

func (h *Handler) CreateAttribute(c *gin.Context) {
	var a CreateAttributeRequest

	err := validator.BindJSON(&a, c.Request)
	if err != nil {
		logrus.WithError(err).Warn("CreateAttribute: invalid request JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	aSrv := a.ToSrv()
	createdSrv, err := h.srv.CreateAttributes(aSrv)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	createdA := Attribute{}
	createdA.FillFromSRV(createdSrv)
	c.JSON(http.StatusCreated, nil)

}
