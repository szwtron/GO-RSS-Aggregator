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

type Feed struct {
	ID        uuid.UUID	`json:"id"`
	Name      string `json:"name"`
	Url       string `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func dbFeedtoFeed(dbFeed db.Feed) Feed {
	return Feed{
		ID: dbFeed.ID,
		Name: dbFeed.Name,
		Url: dbFeed.Url,
		UserID: dbFeed.UserID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
	}
}

func dbFeedstoFeeds(dbFeeds []db.Feed) []Feed {
	feeds := []Feed{}
	for _, dbFeed := range dbFeeds {
		feeds = append(feeds, dbFeedtoFeed(dbFeed))
	}
	return feeds
}

type FeedFollow struct {
	ID        uuid.UUID	`json:"id"`
	FeedID    uuid.UUID `json:"feed_id"`
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func dbFeedFollowtoFeedFollow(dbFeedFollow db.FeedFollow) FeedFollow {
	return FeedFollow{
		ID: dbFeedFollow.ID,
		FeedID: dbFeedFollow.FeedID,
		UserID: dbFeedFollow.UserID,
		CreatedAt: dbFeedFollow.CreatedAt,
		UpdatedAt: dbFeedFollow.UpdatedAt,
	}
}

func dbFeedFollowstoFeedFollows(dbFeedFollows []db.FeedFollow) []FeedFollow {
	FeedFollow := []FeedFollow{}
	for _, dbFeedFollow := range dbFeedFollows {
		FeedFollow = append(FeedFollow, dbFeedFollowtoFeedFollow(dbFeedFollow))
	}
	return FeedFollow
}

type Post struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	Description *string    `json:"description"`
	PublishedAt time.Time `json:"published_at"`
	Url         string    `json:"url"`
	FeedID      uuid.UUID `json:"feed_id"`
}

func dbPosttoPost(dbPost db.Post) Post {
	var description *string
	if dbPost.Description.Valid {
		description = &dbPost.Description.String
	}
	return Post{
		ID: dbPost.ID,
		CreatedAt: dbPost.CreatedAt,
		UpdatedAt: dbPost.UpdatedAt,
		Title: dbPost.Title,
		Description: description,
		PublishedAt: dbPost.PublishedAt,
		Url: dbPost.Url,
		FeedID: dbPost.FeedID,
	}
}

func dbPoststoPosts(dbPosts []db.Post) []Post {
	posts := []Post{}
	for _, dbPost := range dbPosts {
		posts = append(posts, dbPosttoPost(dbPost))
	}
	return posts
}