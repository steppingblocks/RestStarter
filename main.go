package main

import (
	"RestStarter/pkg/posts"
	"RestStarter/pkg/users"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// this Pings the database trying to connect
	db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%v port=%v dbname=%v password=%v username=%v sslmode=disable", os.Getenv("PGHOST"), os.Getenv("PGPORT"), os.Getenv("PGDBNAME"), os.Getenv("PGPASSWORD"), os.Getenv("PGUSER")))
	if err != nil {
		log.Fatalln(err)
	}

	r := gin.Default()

	users.Route(db, r.Group("/"))
	posts.Route(db, r.Group("/"))

	r.Run(fmt.Sprintf("0.0.0.0:%v", os.Getenv("HOST_PORT")))
}
