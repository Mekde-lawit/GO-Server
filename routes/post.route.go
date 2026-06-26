package routes

import (
	"github.com/gorilla/mux"

	"udemygo/services"
)

func CreateRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/post", services.GetAllPosts).Methods("GET")
	
	return router
}