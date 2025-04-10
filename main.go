package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"task/core"
	"task/routes"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}
}

func main() {
	if err := core.LoadDB(os.Getenv("DATABASE_URL")); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer core.DB.Close()

	if err := core.Migrate(core.DB); err != nil {
		log.Fatalf("Failed to run database migrations: %v", err)
	}

	r := mux.NewRouter()

	routes.LoadRoutes(r)

	port := os.Getenv("PORT")
	srv := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf("0.0.0.0:%s", port),
	}

	log.Printf("Starting server on http://0.0.0.0:%s\n", port)
	log.Fatal(srv.ListenAndServe())
}
