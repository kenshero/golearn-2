package routes

import (
	"github.com/gorilla/mux"
)

func GetRoutes() *mux.Router {
	router := mux.NewRouter()
	MainNewsRoutes(router)
	return router
}
