package services

import (
	"go_blog/internal/models"
)

// PostsService represents the service for handling posts-related logic.
type PostsService struct {
	postRepository *models.PostRepository
}

// NewPostsService creates a new instance of the PostsService.
func NewPostsService(pr *models.PostRepository) *PostsService {
	return &PostsService{
		postRepository: pr,
	}
}

// GetAllPosts retrieves all posts from the repository.
func (ps *PostsService) GetAllPosts() ([]models.Post, error) {
	return ps.postRepository.GetAllPosts()
}

// GetPostByID retrieves a post by its ID from the repository.
func (ps *PostsService) GetPostByID(id int) (*models.Post, error) {
	return ps.postRepository.GetPostByID(id)
}
