package main

import (
	"atacado_api_go/controller"
	"atacado_api_go/data"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	data.ConnectDb()

	public := r.Group("/")
	public.GET("/customer", controller.GetCustomers)
	public.GET("/customer/:id", controller.GetCustomerById)
	public.POST("/customer", controller.SaveCustomerWithAddress)
	public.PATCH("/customer/:id", controller.UpdateCustomer)
	public.DELETE("/customer/:id", controller.DeleteCustomer)

	public.GET("/address", controller.GetAddresses)
	public.POST("/address", controller.SaveAddress)

	public.GET("/orders", controller.GetAllOrders)
	public.POST("/order", controller.SaveOrder)

	err := r.Run("localhost:8081")
	if err != nil {
		return
	}
}
