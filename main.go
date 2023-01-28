package main

import (
	"atacado_api_go/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	r := gin.Default()
	public := r.Group("/")
	public.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Hello World"})
	})
	public.GET("/customer", controller.GetCustomers)

	r.Run()
}
