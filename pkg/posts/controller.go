package posts

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"net/http"
)

const (
	GET_POSTS_FROM_USER  = ``
	INSERT_POST_FOR_USER = ``
)

// Post is a struct representing a post in the database
type Post struct {
	ID        string `json:"id" db:"id"`
	CreatedAt string `json:"created_at" db:"created_at"`
	Title     string `json:"title" db:"title"`
	Content   string `json:"content" db:"content"`
	UserID    string `json:"user_id" db:"user_id"`
}

// PostRequest is the request provided to POST /posts/:username
type PostRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func Route(conn *sqlx.DB, router *gin.RouterGroup) {
	rg := router.Group("/posts")

	// GET /posts/:username -> Gets all post for provided user. Returns a 404 if the user does not exist
	// bonus: add pagination
	rg.GET("/:username", func(c *gin.Context) {
		//... YOUR WORK HERE ...//

		c.Status(http.StatusOK)
	})

	// POST /posts/:username -> creates a post for the given username. Returns a 404 if the user does not exist
	// Request body is of the type PostRequest
	rg.POST("/:username", func(c *gin.Context) {
		//... YOUR WORK HERE ...//

		c.Status(http.StatusOK)
	})
}
