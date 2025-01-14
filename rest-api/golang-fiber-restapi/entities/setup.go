package entities

import "github.com/ortizdavid/golang-fiber-restapi/config"

func SetupMigrations() {
	db, _ := config.ConnectDB()
	db.AutoMigrate(&Task{})
	db.AutoMigrate(&User{})
}