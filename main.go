package main

import (
	"github.com/fitnis/order-service/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	api := router.Group("/orders")
	{
		api.POST("", handlers.CreateOrder)
		api.GET("", handlers.GetOrders)
		api.DELETE("/:orderId", handlers.DeleteOrder)
	}

	router.Run(":8083")
}
