package api

import (
	"github.com/gorilla/mux"

	"github.com/alexeyval/test-task-makves-group/internals/app/handlers"
)

func CreateRoutes(userHandler *handlers.UsersHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/get-items", userHandler.List).Methods("GET")

	r.NotFoundHandler = r.NewRoute().HandlerFunc(handlers.NotFound).GetHandler()

	return r
}
