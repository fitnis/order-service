package handlers

import (
	"net/http"

	"github.com/fitnis/order-service/models"
	"github.com/fitnis/order-service/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// curl -X POST -H "Content-Type: application/json" -d '{"patientId":"123","item":"X-ray","priority":"high"}' http://localhost:8083/orders
func CreateOrder(c *gin.Context) {
	var req models.OrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	id := uuid.New().String()
	c.JSON(http.StatusCreated, services.CreateOrder(id, req))
}

// curl http://localhost:8083/orders
func GetOrders(c *gin.Context) {
	c.JSON(http.StatusOK, services.GetOrders())
}

// curl -X DELETE http://localhost:8083/orders/{orderId}
func DeleteOrder(c *gin.Context) {
	id := c.Param("orderId")
	services.DeleteOrder(id)
	c.Status(http.StatusNoContent)
}
