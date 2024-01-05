package db

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Innocent9712/rssagg/internal/database"
	"github.com/Innocent9712/rssagg/internal/models"
	"github.com/google/uuid"
)

func (apiCfg *ApiConfig) HandlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to create feed: %v", err))
		return
	}

	respondWithJSON(w, http.StatusCreated, models.DatabaseFeedToFeed(feed))
}

func (apiCfg *ApiConfig) HandlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to get feeds: %v", err))
		return
	}

	respondWithJSON(w, http.StatusOK, models.DatabaseFeedsToFeeds(feeds))
}
