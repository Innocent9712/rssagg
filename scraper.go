package main

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/Innocent9712/rssagg/internal/database"
	"github.com/google/uuid"
)

// Function to periodical spin up a set of goroutines to fetch rss of a set of feeds
func startScraping(
	db *database.Queries,
	concurency int,
	timeBetweenRequest time.Duration,
) {
	log.Printf("Scraping on %v goroutines every %s duration", concurency, timeBetweenRequest)

	// a sort of setInterval Equivalent in go
	ticker := time.NewTicker(timeBetweenRequest)

	// for range ticker.C { // waits timeBetweenRequest before first run
	for ; ; <-ticker.C { // runs immediately first time
		// do the scraping
		// ...
		feeds, err := db.GetNextFeedsToFetch(
			context.Background(), // global context
			int32(concurency),
		)

		if err != nil {
			log.Printf("Error while getting next feeds to fetch: %v", err)
			continue
		}

		wg := &sync.WaitGroup{} // create a waitGroup to make sure all goroutines finish running before the main loop restarts
		for _, feed := range feeds {
			wg.Add(1) // increment numer of items waitGroup should wait for
			go func(feed database.Feed, wg *sync.WaitGroup) {
				defer wg.Done() // decrement number of items waitGroup should wait for
				// do the scraping
				// ...

				_, err := db.MarkFeedFetched(context.Background(), feed.ID)
				if err != nil {
					log.Printf("Error while marking feed as fetched: %v", err)
					return
				}

				rssFeed, err := urlToFeed(feed.Url)
				if err != nil {
					log.Printf("Error while getting feed: %v", err)
					return
				}

				for _, item := range rssFeed.Channel.Items {
					description := sql.NullString{}

					if item.Description != "" {
						description.String = item.Description
						description.Valid = true
					}

					pubDate, err := time.Parse(time.RFC1123Z, item.PubDate)

					if err != nil {
						log.Printf("Error while parsing pubDate: %v\n", err)
						continue
					}

					// log.Printf("Post found: %v - On feed: %v\n", item.Title, feed.Name)
					_, err = db.CreatePost(context.Background(), database.CreatePostParams{
						ID:          uuid.New(),
						CreatedAt:   time.Now().UTC(),
						UpdatedAt:   time.Now().UTC(),
						Title:       item.Title,
						Description: description,
						Url:         item.Link,
						PublishedAt: pubDate,
						FeedID:      feed.ID,
					})

					if err != nil {
						if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
							continue
						}
						log.Printf("Error while creating post: %v\n", err)
						continue
					}
				}

				log.Printf("Feed %s collected, %v posts found.\n", feed.Name, len(rssFeed.Channel.Items))

			}(feed, wg)
		}
		wg.Wait() // wait for all items to be removed from waitGroup
	}
}
