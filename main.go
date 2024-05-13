package main

import (
	"net/http"

	"github.com/VicenteMza/go_gorm_restapi/db"
	"github.com/VicenteMza/go_gorm_restapi/models"
	"github.com/VicenteMza/go_gorm_restapi/routes"
	"github.com/gorilla/mux"
)

func main() {
	db.DBConnection()
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

	http.ListenAndServe(":3000", r)
}
