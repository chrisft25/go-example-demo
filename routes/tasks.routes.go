package routes

import (
	"github.com/chrisft25/go-example-demo/handlers"
	"github.com/gorilla/mux"
)

func AddTasksRoutes(r *mux.Router) {
    r.HandleFunc("/tasks", handlers.GetTasksHandler).Methods("GET")
	r.HandleFunc("/tasks", handlers.CreateTaskHandler).Methods("POST")
	r.HandleFunc("/tasks/{id}", handlers.GetTaskByIdHandler).Methods("GET")
	r.HandleFunc("/tasks/{id}", handlers.UpdateTaskHandler).Methods("PUT")
	r.HandleFunc("/tasks/{id}", handlers.DeleteTaskHandler).Methods("DELETE")
}