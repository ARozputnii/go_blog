package controllers

// ApplicationController represents the application controller.
type ApplicationController struct {
	HealthCheckController *HealthCheckController
}

// NewApplicationController creates a new instance of the ApplicationController.
func NewApplicationController() *ApplicationController {
	return &ApplicationController{
		HealthCheckController: NewHealthCheckController(),
	}
}
