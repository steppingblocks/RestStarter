package users

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"net/http"
)

const (
	ALL_USERS   = `SELECT id, created_at, first_name, last_name, username FROM "Users"`
	ONE_USER    = `SELECT id, created_at, first_name, last_name, username FROM "Users" WHERE username=$1`
	INSERT_USER = `INSERT INTO "Users" (first_name, last_name, username) VALUES ($1, $2, $3) returning id, first_name, last_name, username, created_at`
)

// User represents a user entry in the database
type User struct {
	ID        string `json:"id" db:"id"`
	CreatedAt string `json:"created_at" db:"created_at"`
	Username  string `json:"username" db:"username"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
}

// UserRequest represents the data required in the POST /users request
type UserRequest struct {
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func Route(conn *sqlx.DB, router *gin.RouterGroup) {
	rg := router.Group("/users")

	// GET /users -> Gets all users from db
	rg.GET("", func(c *gin.Context) {
		users := []User{}
		err := conn.Select(&users, ALL_USERS)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, users)
	})

	// GET /users/:username -> Gets one user from db, returns 404 if not found.
	rg.GET("/:username", func(c *gin.Context) {
		username := c.Param("username")

		user := []User{}
		err := conn.Select(&user, ONE_USER, username)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, err)
			return
		}

		if len(user) != 1 {
			c.Status(http.StatusNotFound)
			return
		}

		c.JSON(http.StatusOK, user[0])
	})

	// POST /users -> creates a new database entry for request body
	rg.POST("/", func(c *gin.Context) {
		var userRequest UserRequest

		if err := c.BindJSON(&userRequest); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, err)
			return
		}

		var user User
		tx := conn.MustBegin()
		err := tx.QueryRowx(INSERT_USER, userRequest.FirstName, userRequest.LastName, userRequest.Username).StructScan(&user)
		if err != nil {
			tx.Rollback()
			c.AbortWithStatusJSON(http.StatusInternalServerError, err)
			return
		}
		err = tx.Commit()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, user)
	})
}
