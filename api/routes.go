package api

import (
	"github.com/alexeyval/test-task-makves-group/internals/app/handlers"
	"github.com/gorilla/mux"
)

func CreateRoutes(userHandler *handlers.UsersHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/get-items", userHandler.List).Methods("GET")

	r.NotFoundHandler = r.NewRoute().HandlerFunc(handlers.NotFound).GetHandler()
	return r
}
