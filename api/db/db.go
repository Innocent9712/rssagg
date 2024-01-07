package db

import (
	"database/sql"
	"log"

	"github.com/Innocent9712/rssagg/internal/database"

	_ "github.com/lib/pq"
)

// api struct
type ApiConfig struct {
	DB *database.Queries
}

func SetupDBConn(dbUrl string) ApiConfig {
	// create new DB connection
	conn, err := sql.Open("postgres", dbUrl)

	if err != nil {
		log.Fatal("Unable to connect to database")
	}

	return ApiConfig{
		DB: database.New(conn),
	}
}
