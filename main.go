package main

import (
	"fmt"
	"net/http"

	"github.com/chrisft25/go-example-demo/db"
	"github.com/chrisft25/go-example-demo/helpers"
	"github.com/chrisft25/go-example-demo/models"
	"github.com/chrisft25/go-example-demo/routes"
)



func main() {

	dbString := helpers.GetEnv("DATABASE_CONNECTION_STRING","")
	db.DBConnection(dbString)

	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	router := routes.NewRouter()

	port := helpers.GetEnv("PORT", "3000")

	http.ListenAndServe(fmt.Sprintf(":%v", port), router)
}