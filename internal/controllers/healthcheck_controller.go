package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// HealthCheckController represents the health check controller.
type HealthCheckController struct{}

// NewHealthCheckController creates a new instance of the HealthCheckController.
func NewHealthCheckController() *HealthCheckController {
	return &HealthCheckController{}
}

// Ping handles the /ping route.
func (h *HealthCheckController) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Service is up and running"})
}
