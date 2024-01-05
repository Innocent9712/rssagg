package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Innocent9712/rssagg/api/db"
	"github.com/Innocent9712/rssagg/internal/routes"
	"github.com/joho/godotenv"
)

func main() {
	// // Testing RSS feed
	// // rss, err := urlToFeed("https://www.theverge.com/rss/index.xml")
	// rss, err := urlToFeed("https://wagslane.dev/index.xml")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// // fmt.Println(rss)
	// for _, item := range rss.Channel.Items {
	// 	fmt.Println(item.Title)
	// }

	godotenv.Load(".env")

	// get the port from env file
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("PORT not found in env")
	}

	dbUrl := os.Getenv("DB_URL")

	if dbUrl == "" {
		log.Fatal("DB_URL not found in env")
	}

	// connect api to db connection
	apiCfg := db.SetupDBConn(dbUrl)

	go startScraping(apiCfg.DB, 10, time.Minute) // start the rss feed scraper

	router := routes.NewRouter(apiCfg)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Println("Starting server on port", port)

	err := srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Port:", port)
}

// go get github.com/joho/godotenv // this installs a package to pull env from .env file
// go mod vendor // generate local copies of dependencies in ./vendor
