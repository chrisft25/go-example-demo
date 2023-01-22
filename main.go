package main

import (
	"net/http"

	"github.com/chrisft25/go-example-demo/db"
	"github.com/chrisft25/go-example-demo/helpers"
	"github.com/chrisft25/go-example-demo/models"
	"github.com/chrisft25/go-example-demo/routes"
)



func main() {

	db.DBConnection(helpers.GetEnv("DATABASE_CONNECTION_STRING"))

	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	router := routes.NewRouter()

	http.ListenAndServe(":3000", router)
}