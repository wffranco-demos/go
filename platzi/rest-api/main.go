package main

import (
	"context"
	"log"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/wffranco-demos/go/platzi/rest-api/handlers"
	"github.com/wffranco-demos/go/platzi/rest-api/server"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")
	JWT_SECRET := os.Getenv("JWT_SECRET")
	DB_URL := os.Getenv("DB_URL")

	s, err := server.NewServer(context.Background(), &server.Config{
		Port:        PORT,
		JWTSecret:   JWT_SECRET,
		DataBaseURL: DB_URL,
	})
	if err != nil {
		log.Fatal("Error creating server: ", err)
	}
	s.Start(RoutesBinder)
}

func RoutesBinder(s server.Server, r *mux.Router) {
	r.HandleFunc("/", handlers.HomeHandler(s)).Methods("GET")
}
