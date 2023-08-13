package main

import (
	"fmt"
	"net/http"

	"github.com/szwtron/rss_aggregator/internal/auth"
	"github.com/szwtron/rss_aggregator/internal/db"
)

type authHandler func(http.ResponseWriter, *http.Request, db.User)

func (apiCfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Auth Error: %v", err))
			return
		}

		user, err := apiCfg.DB.SelectUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error getting user: %v", err))
			return
		}

		handler(w, r, user)
	}
}