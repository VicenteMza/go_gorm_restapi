package main

import (
	"net/http"

	"github.com/VicenteMza/go_gorm_restapi/db"
	"github.com/VicenteMza/go_gorm_restapi/models"
	"github.com/VicenteMza/go_gorm_restapi/routes"
	"github.com/gorilla/mux"
)

const (
	serverPort = "3000"
)

func main() {
	db.DBConnection(serverPort)
	//creo las tablas en la base de datos
	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	r := mux.NewRouter()
	r.HandleFunc("/", routes.HomeHandler)

	//rutas de users
	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", routes.PostUsersHandler).Methods("POST")
	r.HandleFunc("/users/{id}", routes.DeleteUsersHandler).Methods("DELETE")

	//Tasks routes

	r.HandleFunc("/tasks", routes.GetTasksHandler).Methods("GET")
	r.HandleFunc("/tasks/{id}", routes.GetTaskHandler).Methods("GET")
	r.HandleFunc("/tasks", routes.CreateTasksHandler).Methods("POST")
	r.HandleFunc("/tasks/{id}", routes.DeleteTasksHandler).Methods("DELETE")

	http.ListenAndServe(":"+serverPort, r)
}
