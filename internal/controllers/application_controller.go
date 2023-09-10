package controllers

import (
	"go_blog/internal/controllers/api/v1"
	"go_blog/internal/services"
)

// ApplicationController represents the application controller.
type ApplicationController struct {
	HealthCheckController *HealthCheckController
	PostsController       *v1.PostsController
}

// NewApplicationController creates a new instance of the ApplicationController.
func NewApplicationController() *ApplicationController {
	as := services.NewApplicationService()

	return &ApplicationController{
		HealthCheckController: NewHealthCheckController(),
		PostsController:       v1.NewPostsController(as.PostsService),
	}
}
