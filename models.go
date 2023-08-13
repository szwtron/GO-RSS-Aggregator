package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/szwtron/rss_aggregator/internal/db"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Name      string	`json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ApiKey string `json:"api_key"`
}

func dbUsertoUser(dbUser db.User) User {
	return User{
		ID: dbUser.ID,
		Name: dbUser.Name,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		ApiKey: dbUser.ApiKey,
	}
}