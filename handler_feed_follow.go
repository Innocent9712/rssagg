package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	// "github.com/Innocent9712/rssagg/internal/auth"
	"github.com/Innocent9712/rssagg/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID string `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	feedID, err := uuid.Parse(params.FeedID)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid feed ID")
		return
	}

	feedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FeedID:      feedID,
		UserID: user.ID,
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to create feed Follow: %v", err))
		return
	}

	respondWithJSON(w, http.StatusCreated, databaseFeedFollowToFeedFollow(feedFollow))
}

func (apiCfg *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowsIDs, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to get feed follows: %v", err))
		return
	}

	feedFollows := make([]database.Feed, 0)

	for _, feedFollow := range feedFollowsIDs {
		feed, err := apiCfg.DB.GetFeed(r.Context(), feedFollow.FeedID)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to get feed follows: %v", err))
			return
		}

		feedFollows = append(feedFollows, feed)
		
	}

	respondWithJSON(w, http.StatusOK, databaseFeedsToFeeds(feedFollows))
}
