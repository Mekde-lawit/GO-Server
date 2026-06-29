package routes

import (
	"auth/middlewares"
	"auth/services"
	"net/http"

	"github.com/gorilla/mux"
)

func AuthRouter(req *mux.Router) {
	auth := req.PathPrefix("/auth").Subrouter()

	auth.HandleFunc("/login", services.Login).Methods("POST")
	auth.HandleFunc("/register", services.Register).Methods("POST")
	auth.Handle("/me", middlewares.Auth(http.HandlerFunc(services.Me))).Methods("GET")

}
