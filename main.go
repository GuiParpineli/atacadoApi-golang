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

	err := r.Run()
	if err != nil {
		return
	}
}
