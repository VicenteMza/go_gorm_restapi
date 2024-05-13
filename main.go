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

	http.ListenAndServe(":3000", r)
}
