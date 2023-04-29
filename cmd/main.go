package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("error loading .env file: %v", err)
	}

	port := os.Getenv("SERVER_PORT")

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
