package main

import (
	"go_blog/internal/config"
	"go_blog/internal/controllers"
	"os"
)

func main() {
	// load env files
	config.LoadEnv()

	// Initialize the application controller
	ac := controllers.NewApplicationController()

	// Initialize the router with dependencies
	router := config.NewRouter(ac)

	// Initialize and run the router
	server := router.InitRoutes()

	_ = server.Run(":" + os.Getenv("PORT"))
}
