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
