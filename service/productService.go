package service

import (
	"atacado_api_go/data"
	"atacado_api_go/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type createProduct struct {
	Name          string  `json:"name" binding:"required"`
	ProviderName  string  `json:"providerName"`
	Ncm           string  `json:"ncm" `
	Category      string  `json:"category"`
	Inventory     int     `json:"inventory"`
	PurchasePrice float32 `json:"purchasePrice"`
	SalePrice     float32 `json:"salePrice"`
}

func GetProducts(c *gin.Context) {
	var products []models.Product
	data.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{"data": products})
}

func CreateProduct(c *gin.Context) {
	var input createProduct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	product := models.Product{
		Name:          input.Name,
		ProviderName:  input.ProviderName,
		Ncm:           input.Ncm,
		Category:      input.Category,
		Inventory:     input.Inventory,
		PurchasePrice: input.PurchasePrice,
		SalePrice:     input.SalePrice,
	}
	data.DB.Create(&product)
	c.JSON(http.StatusOK, gin.H{"data": product})
}
