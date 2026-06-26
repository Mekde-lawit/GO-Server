package routes

import "github.com/gorilla/mux"

func MainRouter() *mux.Router {
	mainRouter := mux.NewRouter()

	api := mainRouter.PathPrefix("/api").Subrouter()

	PostRouter(api)
	AuthRouter(api)

	return mainRouter
}
