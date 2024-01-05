package main

import (
	"fmt"
	"net/http"

	"github.com/Innocent9712/rssagg/internal/database"
)

func (apiCfg *apiConfig) handlerGetPosts(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := apiCfg.DB.GetPostsOfFeedUserFollows(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to get feeds: %v", err))
		return
	}

	respondWithJSON(w, http.StatusOK, databasePostsToPosts(posts))
}
