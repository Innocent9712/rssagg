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
	Url       string    `json:"url"`
	UserID    string    `json:"user_id"`
}

func databaseFeedToFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID:        dbFeed.ID.String(),
		Name:      dbFeed.Name,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Url:       dbFeed.Url,
		UserID:    dbFeed.UserID.String(),
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
	UserID    string    `json:"user_id"`
	FeedID    string    `json:"feed_id"`
}

type FeedFollows struct {
	FeedFollow
	Feed Feed `json:"feed"`
}

func databaseFeedFollowToFeedFollow(dbFeedFollow database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        dbFeedFollow.ID.String(),
		CreatedAt: dbFeedFollow.CreatedAt,
		UpdatedAt: dbFeedFollow.UpdatedAt,
		UserID:    dbFeedFollow.UserID.String(),
		FeedID:    dbFeedFollow.FeedID.String(),
	}
}

func databaseFeedFollowsToFeedFollows(dbFeedFollows []database.GetFeedFollowsRow) []FeedFollows {
	feedFollows := make([]FeedFollows, 0)
	for _, dbFeedFollow := range dbFeedFollows {
		feedFollow := FeedFollows{
			FeedFollow: FeedFollow{
				ID:        dbFeedFollow.ID.String(), // Assuming ID is part of GetFeedFollowsRow
				CreatedAt: dbFeedFollow.CreatedAt,
				UpdatedAt: dbFeedFollow.UpdatedAt, // Assuming UpdatedAt is part of GetFeedFollowsRow
				UserID:    dbFeedFollow.UserID.String(),
				FeedID:    dbFeedFollow.FeedID.String(),
			},
			Feed: Feed{
				ID:        dbFeedFollow.ID_2.String(), // Assuming FeedID is part of GetFeedFollowsRow
				Name:      dbFeedFollow.Name,          // You might need to fetch this information from the database
				CreatedAt: dbFeedFollow.CreatedAt_2,   // Replace with actual database field
				UpdatedAt: dbFeedFollow.UpdatedAt_2,   // Replace with actual database field
				Url:       dbFeedFollow.Url,           // Replace with actual database field
				UserID:    dbFeedFollow.UserID_2.String(),
			},
		}
		feedFollows = append(feedFollows, feedFollow)
	}
	return feedFollows
}

type Post struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	PublishedAt time.Time `json:"published_at"`
	Url         string    `json:"url"`
	FeedID      string    `json:"feed_id"`
}

func databasePostsToPosts(dbPosts []database.Post) []Post {
	var posts []Post
	for _, dbPost := range dbPosts {

		var description *string
		if dbPost.Description.Valid {
			description = &dbPost.Description.String
		}

		post := Post{
			ID:          dbPost.ID.String(),
			Title:       dbPost.Title,
			Description: description,
			CreatedAt:   dbPost.CreatedAt,
			UpdatedAt:   dbPost.UpdatedAt,
			PublishedAt: dbPost.PublishedAt,
			Url:         dbPost.Url,
			FeedID:      dbPost.FeedID.String(),
		}
		posts = append(posts, post)
	}
	return posts
}
