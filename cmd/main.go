package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	httpConfig "user-services/src/http"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("error loading .env file: %v", err)
	}

	port := os.Getenv("SERVER_PORT")

	readTimeout, err := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))
	writeTimeout, err := strconv.Atoi(os.Getenv("SERVER_WRITE_TIMEOUT"))
	idleTimeout, err := strconv.Atoi(os.Getenv("SERVER_IDLE_TIMEOUT"))

	if err != nil {
		log.Fatalf("error parsing server timeouts: %v", err)
	}

	router := httpConfig.NewRouter()

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		ReadTimeout:  time.Duration(readTimeout) * time.Second,
		WriteTimeout: time.Duration(writeTimeout) * time.Second,
		IdleTimeout:  time.Duration(idleTimeout) * time.Second,
	}

	log.Printf("Server listening on port %s", port)

	err = server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not start server: %v", err)
	}
}
