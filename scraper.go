package main

import (
	"context"
	"log"
	"sync"
	"time"

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
	
func scrapeFeed(db *db.Queries, wg *sync.WaitGroup, feed db.Feed) {
	
	defer wg.Done()

	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
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
		log.Println("Found Item", item.Title, "on feed", feed.Name)
	}
	log.Printf("Feed %s collected, %v posts found", feed.Name, len(rssFeed.Channel.Item))
}