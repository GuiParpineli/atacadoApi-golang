package service

import (
	"atacado_api_go/data"
	"atacado_api_go/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type createCustomerWithAddress struct {
	Name       string `json:"name" binding:"required"`
	Lastname   string `json:"lastname" binding:"required"`
	Email      string `json:"email" binding:"required"`
	Phone      string `json:"phone" binding:"required"`
	Cnpj       string `json:"cnpj" binding:"required"`
	Street     string `json:"street" binding:"required"`
	City       string `json:"city" binding:"required"`
	State      string `json:"state" binding:"required"`
	PostalCode string `json:"postal_code" binding:"required"`
}

type updateCustomerInput struct {
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Cnpj     string `json:"cnpj"`
}

func GetCustomers(c *gin.Context) {
	var customer []models.Customer
	data.DB.Preload("Address").Find(&customer)
	c.JSON(http.StatusOK, gin.H{"data": customer})
}

func GetCustomerById(c *gin.Context) {
	var customer models.Customer

	if err := data.DB.Where("id = ?", c.Param("id")).First(&customer).Preload("Address").Error; err != nil {
		c.JSON(http.StatusNoContent, gin.H{"error": "None customers founded."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": customer})
}

func SaveCustomerWithAddress(c *gin.Context) {
	var input createCustomerWithAddress
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customer := models.Customer{
		Name:     input.Name,
		Lastname: input.Lastname,
		Email:    input.Email,
		Phone:    input.Phone,
		Cnpj:     input.Cnpj,
	}

	data.DB.Create(&customer)

	address := models.Address{
		Street:     input.Street,
		City:       input.City,
		State:      input.State,
		PostalCode: input.PostalCode,
		CustomerID: customer.ID,
	}

	data.DB.Create(&address)

	customer.AddressId = address.ID
	data.DB.Save(&customer)

	c.JSON(http.StatusOK, gin.H{"data": customer})

}

func UpdateCustomer(c *gin.Context) {
	var customer models.Customer
	if err := data.DB.Where("id = ?", c.Param("id")).First(&customer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Customer not found"})
		return
	}

	var input updateCustomerInput
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
