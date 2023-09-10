package controllers

import "go_blog/internal/controllers/api/v1"

// ApplicationController represents the application controller.
type ApplicationController struct {
	HealthCheckController *HealthCheckController
	PostsController       *v1.PostsController
}

// NewApplicationController creates a new instance of the ApplicationController.
func NewApplicationController() *ApplicationController {
	return &ApplicationController{
		HealthCheckController: NewHealthCheckController(),
		PostsController:       v1.NewPostsController(),
	}
}
