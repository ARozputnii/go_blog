package services

import "go_blog/internal/models"

// ApplicationService represents the application-level service.
type ApplicationService struct {
	PostsService *PostsService
}

// NewApplicationService creates a new instance of the ApplicationService.
func NewApplicationService() *ApplicationService {
	ps := NewPostsService(models.NewPostRepository())

	return &ApplicationService{
		PostsService: ps,
	}
}
