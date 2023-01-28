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
	public.POST("/customer", controller.SaveCustomer)
	public.PATCH("/customer/:id", controller.UpdateCustomer)
	public.DELETE("/customer/:id", controller.DeleteCustomer)

	err := r.Run()
	if err != nil {
		return
	}
}
