package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnection(serverPort string) {
	var error error

	var host = "localhost"
	var user = "root"
	var password = "root"
	var dbname = "gorm"

	DSN := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, dbname)

	DB, error = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if error != nil {
		log.Fatal(error)
	} else {
		log.Printf("Server running on port: %s", serverPort)
		log.Printf("URL: http://localhost:%s", serverPort)
		log.Println("DB connected")
	}
}
