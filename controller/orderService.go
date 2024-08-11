package controller

import (
	"atacado_api_go/data"
	"atacado_api_go/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type createProductOrder struct {
	Id       int `json:"id"`
	Quantity int `json:"quantity"`
}

type CreateOrder struct {
	Products   []createProductOrder `json:"products" binding:"required"`
	CustomerID uint                 `json:"customer_id" binding:"required"`
	VendorID   uint                 `json:"vendor_id" binding:"required"`
}

func GetAllOrders(c *gin.Context) {
	var orders []models.ProductOrder
	data.DB.Find(&orders)
	c.JSON(http.StatusOK, gin.H{"data": orders})
}

func SaveOrder(c *gin.Context) {
	var input CreateOrder
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var productsOrder []models.ProductOrder

	for _, productInput := range input.Products {
		var product models.Product
		if err := data.DB.Where("id = ?", productInput.Id).First(&product).Error; err != nil {
			c.JSON(http.StatusNoContent, gin.H{"error": "None products founded."})
			return
		}
		productsOrder = append(productsOrder, models.ProductOrder{Product: product, Quantity: productInput.Quantity})
	}

	var total uint
	for _, product := range productsOrder {
		total += product.Product.SalePrice
	}

	order := models.Order{
		Products:   productsOrder,
		Total:      total,
		CustomerID: input.CustomerID,
		VendorID:   input.VendorID,
	}
	data.DB.Create(&order)
	c.JSON(http.StatusOK, gin.H{"data": order})
}
