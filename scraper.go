package main

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/szwtron/rss_aggregator/internal/db"
)

func startScarping(
	db *db.Queries,
	concurrency int,
	requestInterval time.Duration,
) {
	log.Printf("Scraping on %v goroutines every %s duration", concurrency, requestInterval)
	ticker := time.NewTicker(requestInterval)
	for ; ; <- ticker.C {
		feeds, err := db.GetNextFeedsToFetch(
			context.Background(),
			int32(concurrency),
		)
		if err != nil {
			log.Println("Error fetching feeds: ", err)
			continue
		}

		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)

			go scrapeFeed(db, wg, feed)
		}
		wg.Wait()
	}
}
	
func scrapeFeed(database *db.Queries, wg *sync.WaitGroup, feed db.Feed) {
	
	defer wg.Done()

	_, err := database.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Println("Error making feed as fetched:", err)
		return
	}

	rssFeed, err := urlToFeed(feed.Url)
	if err != nil {
		log.Println("Error fetching feed:", err)
		return
	}

	for _, item := range rssFeed.Channel.Item {
		description := sql.NullString{}
		if item.Description != "" {
			description.String = item.Description
			description.Valid = true
		}

		publishedAt, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			log.Printf("Error parsing published date %v with error %v:", item.PubDate, err)
		} 

		_, err = database.CreatePost(context.Background(), db.CreatePostParams{
			ID: uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Title: item.Title,
			Description: description,
			PublishedAt: publishedAt,
			Url: item.Link,
			FeedID: feed.ID,
		})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key") {
				continue
			}
			log.Println("Error creating post:", err)
		}
	}


	log.Printf("Feed %s collected, %v posts found", feed.Name, len(rssFeed.Channel.Item))
}