package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/szwtron/rss_aggregator/internal/db"
)

func (apiCfg *apiConfig)handlerCreateFeed(w http.ResponseWriter, r *http.Request, user db.User) {
	type parameters	struct {
		Name string `json:name`
		URL string `json:url`
	}

	decoder := json.NewDecoder(r.Body)

	params := &parameters{}
	err := decoder.Decode(params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Error parsing JSON")
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), db.CreateFeedParams{
		ID:  uuid.New(),
		Name: params.Name,
		Url: params.URL,
		UserID: user.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error creating feed: %v", err))
		return
	}

	respondWithJSON(w, 201, dbFeedtoFeed(feed))
}

func (apiCfg *apiConfig)handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.SelectAllFeeds(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Can not get feeds: %v", err))
		return
	}

	respondWithJSON(w, 200, dbFeedstoFeeds(feeds))
}