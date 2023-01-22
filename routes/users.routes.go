package routes

import (
	"encoding/json"
	"net/http"

	"github.com/chrisft25/go-example-demo/db"
	"github.com/chrisft25/go-example-demo/models"
	"github.com/gorilla/mux"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request){
	var users []models.User

	db.DB.Find(&users)
	json.NewEncoder(w).Encode(&users)

}

func GetUserByIdHandler(w http.ResponseWriter, r *http.Request){
	var user models.User
	params := mux.Vars(r)

	db.DB.First(&user, params["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound) // 404
		w.Write([]byte("User not found"))
		return
	}

	db.DB.Model(&user).Association("Tasks").Find(&user.Tasks)

	json.NewEncoder(w).Encode(&user)
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request){
	var user models.User

	json.NewDecoder(r.Body).Decode(&user)

	createdUser := db.DB.Create(&user)
	err := createdUser.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(&user)
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request){
	var user models.User
	params := mux.Vars(r)

	db.DB.First(&user, params["id"])

	
	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound) // 404
		w.Write([]byte("User not found"))
		return
	}
	
	json.NewDecoder(r.Body).Decode(&user)

	db.DB.Save(&user)

	json.NewEncoder(w).Encode(&user)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request){
	var user models.User
	params := mux.Vars(r)

	db.DB.First(&user, params["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound) // 404
		w.Write([]byte("User not found"))
		return
	}

	db.DB.Delete(&user)

	w.WriteHeader(http.StatusOK)
}
