package routes

import (
	"github.com/gorilla/mux"
	"auth/services"
)

func AuthRouter(req *mux.Router) {
	auth := req.PathPrefix("/auth").Subrouter()

	auth.HandleFunc("/login", services.Login).Methods("POST")
	auth.HandleFunc("/register", services.Register).Methods("POST")
	auth.HandleFunc("/logout", services.Logout).Methods("POST")
	auth.HandleFunc("/protected", services.Protected).Methods("GET")

}