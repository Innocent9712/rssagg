package db

import "net/http"

func HandlerError(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 400, "something went wrong")
}
