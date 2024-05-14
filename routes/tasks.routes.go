package routes

import (
	"encoding/json"
	"net/http"

	"github.com/VicenteMza/go_gorm_restapi/db"
	"github.com/VicenteMza/go_gorm_restapi/models"
	"github.com/gorilla/mux"
)

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task // task type
	db.DB.Find(&tasks)
	json.NewEncoder(w).Encode(&tasks)
}

func CreateTasksHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)
	createdTask := db.DB.Create(&task)
	err := createdTask.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(&task)
}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task // task type
	param := mux.Vars(r)

	db.DB.First(&task, param["id"])
	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound) // 404
		w.Write([]byte("Task not found"))
		return
	}

	json.NewEncoder(w).Encode(&task)
}

func DeleteTasksHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	params := mux.Vars(r)
	db.DB.First(&task, params["id"])
	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound) // 404
		w.Write([]byte("Task not found"))
		return
	}
	//borrar fisicamente el dato de la tabla
	db.DB.Unscoped().Delete(&task)
	w.WriteHeader(http.StatusNoContent) // 204
}
