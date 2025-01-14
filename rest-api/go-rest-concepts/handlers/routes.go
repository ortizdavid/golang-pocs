package handlers

import (
	"net/http"
	"gorm.io/gorm"
)

func RegisterRoutes(router* http.ServeMux, db *gorm.DB) {
	NewTaskHandler(db).Routes(router)
	NewUserHandler(db).Routes(router)
	NewReportHandler(db).Routes(router)
}