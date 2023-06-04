package main

import (
	"database/sql"
	"fmt"
	"golang-sql/internal/database"
	"log"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

type CustomContext struct {
	echo.Context
	apiConfig
}

var e = echo.New()

func main() {
	portString := os.Getenv("PORT")
	fmt.Printf("The port is: %s", portString)
	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not found in the environment")
	}

	conn, err := sql.Open("postgres", dbURL)

	if err != nil {
		log.Fatal("Can't connect to database:", err)
	}

	db := database.New(conn)
	apiCfg := apiConfig{
		DB: db,
	}

	go startScraping(db, 10, time.Minute)

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &CustomContext{c, apiCfg}
			// fmt.Printf("CustomContext instance: %+v", cc)
			return next(cc)
		}
	})

	e.POST("/users", handlerCreateUser)
	e.GET("/users", middlewareAuth(handlerGetUserByAPIKey))

	e.GET("/feeds", handlerGetFeeds)
	e.POST("/feeds", middlewareAuth(handlerCreateFeed))

	e.GET("/feed_follows", middlewareAuth(handlerGetFollowFeeds))
	e.POST("/feed_follows", middlewareAuth(handlerCreateFeedFollow))
	e.DELETE("/feed_follows/:feed_follow_id", middlewareAuth(handlerDeleteFeedFollow))

	e.GET("/posts", middlewareAuth(handlerGetPostsForUser))

	e.Logger.Printf(fmt.Sprintf("Listening on port %s", portString))
	e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%s", portString)))
}
