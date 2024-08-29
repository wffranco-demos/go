package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// load .env file
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}
	fmt.Println("PORT is", os.Getenv("PORT"))
}
