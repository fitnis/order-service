package main

import (
	"github.com/fitnis/order-service/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	api := router.Group("/orders")
	{
		handlers.RegisterOrderRoutes(api)
	}

	router.Run(":8083")
}
