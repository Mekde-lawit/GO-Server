package routes

import (
	"github.com/gorilla/mux"

	"udemygo/services"
)

func CreateRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/post", services.GetAllPosts).Methods("GET")
	router.HandleFunc("/post", services.CreatePost).Methods("POST")
	router.HandleFunc("/post/{id}", services.GetPost).Methods("GET")
	router.HandleFunc("/post/{id}", services.UpdatePost).Methods("PUT")
	router.HandleFunc("/post/{id}", services.DeletePost).Methods("DELETE")

	return router
}
