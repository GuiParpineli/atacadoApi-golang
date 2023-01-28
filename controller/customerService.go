package controller

import (
	"atacado_api_go/data"
	"atacado_api_go/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCustomers(c *gin.Context) {
	var customer []models.Customer
	data.DB.Find(&customer)
	c.JSON(http.StatusOK, gin.H{"data": customer})
}
