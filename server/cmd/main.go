package main

import (
	"fmt"
	"os"

	"server/db"
	"server/internal/user"
	ws "server/internal/websocket"
	"server/router"

	"github.com/joho/godotenv"
)


func main() {
	
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}

	// Get the Postgres URL from environment variables
	postgresURL := fmt.Sprintf("postgres://%s:%s@localhost:%s/%s?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
	)
	if postgresURL == "" {
		fmt.Println("POSTGRES_URL is not set in the environment")
		return
	}

	// Initialize the database connection
	db, err := db.NewDatabase(postgresURL)
	if err != nil {
		fmt.Printf("Error connecting to database: %v\n", err)
		return
	}
	userRep := user.NewRepository(db.GetDB())
	userSvc := user.NewService(userRep)
	userHandler := user.NewHandler(userSvc)

	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)
	go hub.Run()

	router.InitRouter(userHandler, wsHandler)
	router.Start("localhost:8080")
}