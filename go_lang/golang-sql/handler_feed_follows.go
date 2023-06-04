package main

import (
	"fmt"
	"golang-sql/internal/database"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func handlerCreateFeedFollow(cc *CustomContext, user database.User) error {
	type body struct {
		FeedId uuid.UUID `json:"feed_id"`
	}

	var reqBody body

	if err := cc.Bind(&reqBody); err != nil {
		return cc.JSON(http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %s", err))
	}

	feed_follow, err := cc.DB.CreateFeedFollow(cc.Request().Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		UserID:    user.ID,
		FeedID:    reqBody.FeedId,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})

	if err != nil {
		return cc.JSON(http.StatusBadRequest, fmt.Sprintf("Something happened creating the feed: %s", err))
	}

	return cc.JSON(http.StatusCreated, databaseFeedFollowToFeedFollow(feed_follow))
}

func handlerGetFollowFeeds(cc *CustomContext, user database.User) error {

	feeds, err := cc.DB.GetFeedFollows(cc.Request().Context(), user.ID)

	if err != nil {
		return cc.JSON(http.StatusInternalServerError, fmt.Sprintf("Something bad happened getting trying to get your feeds: %s", err))
	}

	return cc.JSON(http.StatusOK, databaseFeedsFollowToFeedsFollow(feeds))
}

func handlerDeleteFeedFollow(cc *CustomContext, user database.User) error {

	feedFollowId, err := uuid.Parse(cc.Param("feed_follow_id"))

	if err != nil {
		return cc.JSON(http.StatusBadRequest, fmt.Sprint("Enter a proper UUID: %v", err))
	}

	if err := cc.DB.DeleteFeedFollow(cc.Request().Context(), database.DeleteFeedFollowParams{
		ID:     feedFollowId,
		UserID: user.ID,
	}); err != nil {
		return cc.JSON(http.StatusInternalServerError, fmt.Sprintf("Something bad happened getting trying to get your feeds: %s", err))
	}

	return cc.JSON(http.StatusOK, struct{}{})
}
