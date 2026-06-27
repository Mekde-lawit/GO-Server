package routes

import (
	"udemygo/services"

	"github.com/gorilla/mux"
)

func PostRouter(res *mux.Router) {
	posts := res.PathPrefix("/posts").Subrouter()

	posts.HandleFunc("", services.GetAllPosts).Methods("GET")
	posts.HandleFunc("", services.CreatePost).Methods("POST")
	posts.HandleFunc("/{id}", services.GetPost).Methods("GET")
	posts.HandleFunc("/{id}", services.UpdatePost).Methods("PUT")
	posts.HandleFunc("/{id}", services.DeletePost).Methods("DELETE")
}
