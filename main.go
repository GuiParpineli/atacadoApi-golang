package main

import (
	"atacado_api_go/data"
	"atacado_api_go/models"
	"atacado_api_go/service"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	data.ConnectDb()
	data.DB.LogMode(true)
	data.DB.AutoMigrate(&models.Product{}, &models.Address{}, &models.Customer{}, &models.Order{}, &models.ProductOrder{})

	public := r.Group("/")
	public.GET("/customer", service.GetCustomers)
	public.GET("/customer/:id", service.GetCustomerById)
	public.POST("/customer", service.SaveCustomerWithAddress)
	public.PATCH("/customer/:id", service.UpdateCustomer)
	public.DELETE("/customer/:id", service.DeleteCustomer)

	public.GET("/products", service.GetProducts)
	public.POST("/products", service.CreateProduct)

	public.GET("/address", service.GetAddresses)
	public.POST("/address", service.SaveAddress)

	public.GET("/orders", service.GetAllOrders)
	public.GET("/orders/:id", service.GetOrderById)
	public.POST("/order", service.SaveOrder)

	err := r.Run("localhost:8081")
	if err != nil {
		return
	}
}
