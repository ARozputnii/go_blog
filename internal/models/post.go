package models

// Post represents the structure of a post.
type Post struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// PostRepository represents a repository for managing posts.
type PostRepository struct {
	posts []Post
}

// NewPostRepository creates a new instance of the PostRepository.
func NewPostRepository() *PostRepository {
	// tmp mocked data
	return &PostRepository{
		posts: []Post{
			{ID: 1, Title: "Post 1", Content: "Content for Post 1"},
			{ID: 2, Title: "Post 2", Content: "Content for Post 2"},
		},
	}
}

// GetAllPosts retrieves all posts from DB.
func (pr *PostRepository) GetAllPosts() ([]Post, error) {
	return pr.posts, nil
}

// GetPostByID retrieves a post by its ID from DB.
func (pr *PostRepository) GetPostByID(id int) (*Post, error) {
	for _, post := range pr.posts {
		if post.ID == id {
			return &post, nil
		}
	}
	return nil, nil // Return nil if not found
}
