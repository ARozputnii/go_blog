// internal/config/routes.go

package config

import (
	"github.com/gin-gonic/gin"
	"go_blog/internal/controllers"
)

// Router represents the application router.
type Router struct {
	AC *controllers.ApplicationController
	// ... DB in coming
}

// NewRouter creates a new instance of the Router.
func NewRouter(ac *controllers.ApplicationController) *Router {
	return &Router{
		AC: ac,
	}
}

// InitRoutes initializes and sets up the application routes.
func (r *Router) InitRoutes() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", r.AC.HealthCheckController.Ping)

	return router
}
