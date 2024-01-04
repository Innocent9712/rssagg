package main

import (
	"time"

	"github.com/Innocent9712/rssagg/internal/database"
	// "github.com/google/uuid"
)

// Trying to convert the keys of User to a different format than what's in the db
type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ApiKey    string    `json:"api_key"`
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID.String(),
		ApiKey:    dbUser.ApiKey,
		Name:      dbUser.Name,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
	}
}

type Feed struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Url	string	`json:"url"`
	UserID string `json:"user_id"`
}


func databaseFeedToFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID:        dbFeed.ID.String(),
		Name:      dbFeed.Name,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Url:	dbFeed.Url,
		UserID: dbFeed.UserID.String(),
	}
}

func databaseFeedsToFeeds(dbFeeds []database.Feed) []Feed {
	feeds := make([]Feed, 0)
	for _, feed := range dbFeeds {
		feeds = append(feeds, databaseFeedToFeed(feed))
	}
	return feeds
}


type FeedFollow struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID   string `json:"user_id"`
	FeedID   string `json:"feed_id"`
}

func databaseFeedFollowToFeedFollow(dbFeedFollow database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        dbFeedFollow.ID.String(),
		CreatedAt: dbFeedFollow.CreatedAt,
		UpdatedAt: dbFeedFollow.UpdatedAt,
		UserID:   dbFeedFollow.UserID.String(),
		FeedID:   dbFeedFollow.FeedID.String(),
	}
}