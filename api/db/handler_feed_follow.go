package db

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	// "github.com/Innocent9712/rssagg/internal/auth"
	"github.com/Innocent9712/rssagg/internal/database"
	"github.com/Innocent9712/rssagg/internal/models"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (apiCfg *ApiConfig) HandlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
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
		FeedID:    feedID,
		UserID:    user.ID,
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to create feed Follow: %v", err))
		return
	}

	respondWithJSON(w, http.StatusCreated, models.DatabaseFeedFollowToFeedFollow(feedFollow))
}

func (apiCfg *ApiConfig) HandlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollows, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to get feed follows: %v", err))
		return
	}

	respondWithJSON(w, http.StatusOK, models.DatabaseFeedFollowsToFeedFollows(feedFollows))
}

func (apiCfg *ApiConfig) HandlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowID, err := uuid.Parse(chi.URLParam(r, "feedFollowID"))

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid feed follow ID")
		return
	}

	err = apiCfg.DB.DeleteFeedFollows(r.Context(), database.DeleteFeedFollowsParams{
		ID:     feedFollowID,
		UserID: user.ID,
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to delete feed follow: %v", err))
		return
	}

	respondWithJSON(w, http.StatusOK, nil)
}
