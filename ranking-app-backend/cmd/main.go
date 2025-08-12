package main

import (
	"log"
	"net/http"

	"ranking-app-backend/internal/database"
	"ranking-app-backend/internal/routes"
)

func main() {
	// Initialize the database
	db, err := database.InitDB("./ranking.db")
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// Create tables
	if err := database.CreateTables(db); err != nil {
		log.Fatalf("Failed to create tables: %v", err)
	}

	// Setup routes and start the server
	router := routes.SetupRoutes(db)
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
