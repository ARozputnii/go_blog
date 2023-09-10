package main

import (
	"go_blog/internal/config"
	"go_blog/internal/controllers"
)

func main() {
	// Load environment variables into the configuration structure
	env, err := config.LoadEnv()
	if err != nil {
		panic("Failed to load environment variables")
	}

	// Initialize the application controller
	ac := controllers.NewApplicationController()

	// Initialize the router with dependencies
	router := config.NewRouter(ac)

	// Initialize and run the router
	server := router.InitRoutes()

	_ = server.Run(":" + env.Port)
}
