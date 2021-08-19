package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// HealthCheck GET /healthcheck
func HealthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	}
}
