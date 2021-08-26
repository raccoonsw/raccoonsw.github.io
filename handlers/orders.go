package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"restApiProject/database"
	"restApiProject/grpc_email_client"
	"restApiProject/validation"
)

// CreateOrder POST /orders
func CreateOrder(db database.OrdersInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		receivedOrder, err := validation.ValidateJsonOrder(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		order, err := db.CreateOrder(receivedOrder)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}
		err = grpc_email_client.Client(order)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"id": order.Id})
	}
}
