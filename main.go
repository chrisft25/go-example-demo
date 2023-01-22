package main

import (
	"net/http"

	"github.com/chrisft25/go-example-demo/db"
	"github.com/chrisft25/go-example-demo/helpers"
	"github.com/chrisft25/go-example-demo/models"
	"github.com/chrisft25/go-example-demo/routes"
	"github.com/gorilla/mux"
)



func main() {

	db.DBConnection(helpers.GetEnv("DATABASE_CONNECTION_STRING"))

	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	router := mux.NewRouter()

	router.HandleFunc("/", routes.HomeHandler)

	// Users routes
	router.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	router.HandleFunc("/users", routes.CreateUserHandler).Methods("POST")
	router.HandleFunc("/users/{id}", routes.GetUserByIdHandler).Methods("GET")
	router.HandleFunc("/users/{id}", routes.UpdateUserHandler).Methods("PUT")
	router.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")

	// Tasks routes
	router.HandleFunc("/tasks", routes.GetTasksHandler).Methods("GET")
	router.HandleFunc("/tasks", routes.CreateTaskHandler).Methods("POST")
	router.HandleFunc("/tasks/{id}", routes.GetTaskByIdHandler).Methods("GET")
	router.HandleFunc("/tasks/{id}", routes.UpdateTaskHandler).Methods("PUT")
	router.HandleFunc("/tasks/{id}", routes.DeleteTaskHandler).Methods("DELETE")

	http.ListenAndServe(":3000", router)
}