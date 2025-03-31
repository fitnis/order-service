package main

import (
	"github.com/fitnis/order-service/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	api := router.Group("/orders")
	{
		api.POST("/new", handlers.CreateOrder)
		api.GET("/get", handlers.GetOrders)
		api.DELETE("/delete/:orderId", handlers.DeleteOrder)
	}

	router.Run(":8083")
}
