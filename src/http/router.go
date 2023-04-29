package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() http.Handler {
	r := mux.NewRouter()

	// Users routes
	r.HandleFunc("/users", CreateUserHandler).Methods("POST")
	r.HandleFunc("/users", GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id:[0-9]+}", GetUserHandler).Methods("GET")
	r.HandleFunc("/users/{id:[0-9]+}", UpdateUserHandler).Methods("PUT")
	r.HandleFunc("/users/{id:[0-9]+}", DeleteUserHandler).Methods("DELETE")

	// Health check
	r.HandleFunc("/health", HealthCheckHandler).Methods("GET")

	return r
}
