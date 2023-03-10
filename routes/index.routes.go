package routes

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
    r := mux.NewRouter()
    AddUsersRoutes(r)
	AddTasksRoutes(r)
    return r
}