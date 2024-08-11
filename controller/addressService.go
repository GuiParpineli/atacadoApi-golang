package controller

import (
	"atacado_api_go/data"
	"atacado_api_go/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateAddress struct {
	Street     string `json:"street" binding:"required"`
	City       string `json:"city" binding:"required"`
	State      string `json:"state" binding:"required"`
	PostalCode string `json:"postal_code" binding:"required"`
}

func GetAddresses(c *gin.Context) {
	var address []models.Address
	data.DB.Find(&address)
	c.JSON(http.StatusOK, gin.H{"data": address})
}

func SaveAddress(c *gin.Context) {
	var input CreateAddress
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	address := models.Address{State: input.Street, City: input.City, Street: input.Street, PostalCode: input.PostalCode}
	data.DB.Create(&address)
	c.JSON(http.StatusOK, gin.H{"data": address})
}
