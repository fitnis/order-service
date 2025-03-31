package services

import (
	"github.com/fitnis/order-service/models"
)

var orders = make(map[string]models.OrderRequest)

func CreateOrder(id string, req models.OrderRequest) models.GenericResponse {
	orders[id] = req
	return models.GenericResponse{Message: "Order created"}
}

func GetOrders() []models.OrderRequest {
	var list []models.OrderRequest
	for _, o := range orders {
		list = append(list, o)
	}
	return list
}

func DeleteOrder(id string) {
	delete(orders, id)
}
