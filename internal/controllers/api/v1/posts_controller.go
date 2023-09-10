package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// PostsController represents the posts controller.
type PostsController struct{}

// NewPostsController creates a new instance of the PostsController.
func NewPostsController() *PostsController {
	return &PostsController{}
}

// GetAllPosts handles the GET request to fetch all posts.
func (pc *PostsController) GetAllPosts(c *gin.Context) {
	posts := []string{"Post 1", "Post 2", "Post 3"}

	c.JSON(http.StatusOK, posts)
}

// GetPostByID handles the GET request to fetch a post by its ID.
func (pc *PostsController) GetPostByID(c *gin.Context) {
	postID := c.Param("id")
	post := "Post " + postID

	c.JSON(http.StatusOK, post)
}
