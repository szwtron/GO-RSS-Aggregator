package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/szwtron/rss_aggregator/internal/db"
)

func (apiCfg *apiConfig)handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user db.User) {
	type parameters	struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := &parameters{}
	err := decoder.Decode(params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Error parsing JSON")
		return
	}

	feedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), db.CreateFeedFollowParams{
		ID:  uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: user.ID,
		FeedID: params.FeedID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error creating feed follow: %v", err))
		return
	}

	respondWithJSON(w, 201, dbFeedFollowtoFeedFollow(feedFollow))
}

func (apiCfg *apiConfig)handlerGetFeedFollow(w http.ResponseWriter, r *http.Request, user db.User) {
	feedFollow, err := apiCfg.DB.GetFeedFollow(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error getting feed follow: %v", err))
		return
	}

	respondWithJSON(w, 201, dbFeedFollowstoFeedFollows(feedFollow))
}

func (apiCfg *apiConfig)handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user db.User) {
	feedFollowIDStr := chi.URLParam(r, "feedFollowID")
	feedFollowID, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Could not parse feed follow ID: %v", err))
		return
	}

	err = apiCfg.DB.DeleteFeedFollow(r.Context(), db.DeleteFeedFollowParams{
		ID: feedFollowID,
		UserID: user.ID,
	})
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Could not delete feed follow: %v", err))
		return
	}

	respondWithJSON(w, 200, nil)
}
