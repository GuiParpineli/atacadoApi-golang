package controller

import (
	"atacado_api_go/data"
	"atacado_api_go/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateCustomer struct {
	Name     string `json:"name" binding:"required"`
	Lastname string `json:"lastname" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Cpf      string `json:"cpf" binding:"required"`
}
type UpdateCustomerInput struct {
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Cpf      string `json:"cpf"`
}

func GetCustomers(c *gin.Context) {
	var customer []models.Customer
	data.DB.Find(&customer)
	c.JSON(http.StatusOK, gin.H{"data": customer})
}

func GetCustomerById(c *gin.Context) {
	var customer models.Customer

	if err := data.DB.Where("id = ?", c.Param("id")).First(&customer).Preload("AddressId").Error; err != nil {
		c.JSON(http.StatusNoContent, gin.H{"error": "None customers founded."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": customer})
}

func SaveCustomer(c *gin.Context) {
	var input CreateCustomer
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	customer := models.Customer{Name: input.Name, Lastname: input.Lastname, Email: input.Email,
		Phone: input.Phone, Cpf: input.Cpf}
	data.DB.Create(&customer)
	c.JSON(http.StatusOK, gin.H{"data": customer})
}

func UpdateCustomer(c *gin.Context) {
	var customer models.Customer
	if err := data.DB.Where("id = ?", c.Param("id")).First(&customer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Customer not found"})
		return
	}

	var input UpdateCustomerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data.DB.Model(&customer).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": customer})
}

func DeleteCustomer(c *gin.Context) {
	var customer models.Customer
	if err := data.DB.Where("id = ?", c.Param("id")).First(&customer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Customer not found"})
		return
	}
	data.DB.Delete(&customer)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
