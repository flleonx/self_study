package main

import (
	"fmt"
	"golang-sql/internal/database"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func handlerCreateFeed(cc *CustomContext, user database.User) error {
	type body struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	var reqBody body

	if err := cc.Bind(&reqBody); err != nil {
		return cc.JSON(http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %s", err))
	}

	feed, err := cc.DB.CreateFeed(cc.Request().Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		UserID:    user.ID,
		Name:      reqBody.Name,
		Url:       reqBody.Url,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})

	if err != nil {
		return cc.JSON(http.StatusBadRequest, fmt.Sprintf("Something happened creating the feed: %s", err))
	}

	return cc.JSON(http.StatusCreated, feed)
}

func handlerGetFeeds(c echo.Context) error {
	cc := c.(*CustomContext)

	feeds, err := cc.DB.GetFeeds(cc.Request().Context())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Something bad happened getting trying to get the feeds: %s", err))
	}

	return c.JSON(http.StatusOK, databaseFeedsToFeeds(feeds))
}
