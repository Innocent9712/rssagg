package db

import (
	"fmt"
	"net/http"

	"github.com/Innocent9712/rssagg/internal/auth"
	"github.com/Innocent9712/rssagg/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *ApiConfig) MiddlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), r.Header.Get("X-API-Key"))
		apiKey, err := auth.GetAPIKey(r.Header)

		if err != nil {
			respondWithError(w, http.StatusUnauthorized, fmt.Sprintf("Auth error: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)

		if err != nil {
			respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to get user: %v", err))
			return
		}

		handler(w, r, user)
	}
}
