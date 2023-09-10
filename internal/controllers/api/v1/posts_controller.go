package v1

import (
	"github.com/gin-gonic/gin"
	"go_blog/internal/services"
	"net/http"
	"strconv"
)

// PostsController represents the controller for handling posts-related routes.
type PostsController struct {
	postsService *services.PostsService
}

// NewPostsController creates a new instance of the PostsController.
func NewPostsController(ps *services.PostsService) *PostsController {
	return &PostsController{
		postsService: ps,
	}
}

// GetAllPosts handles GET /api/v1/posts route.
func (pc *PostsController) GetAllPosts(c *gin.Context) {
	posts, err := pc.postsService.GetAllPosts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, posts)
}

// GetPostByID handles GET /api/v1/posts/:id route.
func (pc *PostsController) GetPostByID(c *gin.Context) {
	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	post, err := pc.postsService.GetPostByID(postID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	if post == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(http.StatusOK, post)
}
