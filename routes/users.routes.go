package routes

import (
	"encoding/json"
	"net/http"

	"github.com/VicenteMza/go_gorm_restapi/db"
	"github.com/VicenteMza/go_gorm_restapi/models"
	"github.com/gorilla/mux"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.DB.Find(&users)
	json.NewEncoder(w).Encode(&users)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User // user type
	params := mux.Vars(r)
	db.DB.First(&user, params["id"])
	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound) // 404
		w.Write([]byte("User not found"))
	}

	//Asociasion de la tarea con el usuario
	db.DB.Model(&user).Association("Tasks").Find(&user.Tasks)

	json.NewEncoder(w).Encode(&user)
}

func PostUsersHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	createUser := db.DB.Create(&user)
	err := createUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&user)
}

func DeleteUsersHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)
	db.DB.First(&user, params["id"])
	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound) // 404
		w.Write([]byte("User not found"))
		return
	}
	/*
	   //No elimina fisicamente de la base de datos
	   	db.DB.Delete(&user)
	   	w.WriteHeader(http.StatusOK)

	*/
	// Elimina fisicamente de la base de datos
	db.DB.Unscoped().Delete(&user)
	w.WriteHeader(http.StatusOK)
}
