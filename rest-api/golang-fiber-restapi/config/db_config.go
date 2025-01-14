package config

import (
	"fmt"
	"os"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	LoadDotEnv()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", 
		os.Getenv("DB_USER"), 
		os.Getenv("DB_PASSWORD"), 
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"), 
		os.Getenv("DB_NAME"),
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect DB")
	}
	return db, nil
}

func DisconnectDB(db *gorm.DB) {
	conn, err := db.DB()
	if err != nil {
		panic("Failed to disconnect DB")
	}
	conn.Close()
}