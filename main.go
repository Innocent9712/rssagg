package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Innocent9712/rssagg/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

// api struct
type apiConfig struct {
	DB *database.Queries
}

func main() {
	fmt.Println("Hi There")
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

	// create new DB connection
	conn, err := sql.Open("postgres", dbUrl)

	if err != nil {
		log.Fatal("Unable to connect to database")
	}

	// connect api to db connection
	apiCfg := apiConfig{
		DB: database.New(conn),
	}

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerError)
	v1Router.Post("/users", apiCfg.handlerCreateUser)
	v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.handlerGetUser))
	v1Router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Println("Starting server on port", port)

	err = srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Port:", port)
}

// go get github.com/joho/godotenv // this installs a package to pull env from .env file
// go mod vendor // generate local copies of dependencies in ./vendor
