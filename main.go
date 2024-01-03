package main

import (
	"fmt"
	"log"
	"os"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func main()  {
	fmt.Println("Hi There")
	godotenv.Load(".env")


	port := os.Getenv("PORT")


	if port == "" {
		log.Fatal("PORT not found in env")
	}

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge: 300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerError)

	router.Mount("/v1", v1Router)



	srv := &http.Server{
		Handler: router,
		Addr: ":" + port,
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