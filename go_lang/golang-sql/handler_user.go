package main

import (
	"fmt"
	"golang-sql/internal/database"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// func (cfg *apiConfig) createUser(cc echo.Context) error {
func handlerCreateUser(c echo.Context) error {
	cc := c.(*CustomContext)

	type body struct {
		Name string `json:"name"`
	}

	var reqBody body

	if err := cc.Bind(&reqBody); err != nil {
		return cc.JSON(http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %s", err))
	}

	user, err := cc.DB.CreateUser(cc.Request().Context(), database.CreateUserParams{
		ID:        uuid.New(),
		Name:      reqBody.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})

	if err != nil {
		return cc.JSON(http.StatusBadRequest, fmt.Sprintf("Could not create user %s", err))
	}

	return cc.JSON(http.StatusCreated, databaseUsertoUser(user))
}

func handlerGetUserByAPIKey(cc *CustomContext, user database.User) error {
	return cc.JSON(http.StatusOK, databaseUsertoUser(user))
}

func handlerGetPostsForUser(cc *CustomContext, user database.User) error {
	posts, err := cc.DB.GetPostForUser(cc.Request().Context(), database.GetPostForUserParams{
		UserID: user.ID,
		Limit:  10,
	})

	if err != nil {
		return cc.JSON(http.StatusBadRequest, fmt.Sprintf("Couldn't get posts: %s", err))
	}

	return cc.JSON(http.StatusOK, databasePostsToPosts(posts))
}
