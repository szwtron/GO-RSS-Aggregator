// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package db

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Feed struct {
	ID          uuid.UUID
	Name        string
	Url         string
	UserID      uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
	LastFetchAt sql.NullTime
}

type FeedFollow struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    uuid.UUID
	FeedID    uuid.UUID
}

type User struct {
	ID        uuid.UUID
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	ApiKey    string
}
