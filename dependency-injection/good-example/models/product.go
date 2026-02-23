package models

import "time"

type Product struct {
    ID          uint      `gorm:"primaryKey" json:"id"`
    Name        string    `gorm:"size:100;not null" json:"name"`
    Price       float64   `gorm:"not null" json:"price"`
    Stock       int       `gorm:"default:0" json:"stock"`
    Description string    `gorm:"type:text" json:"description"`
    CreatedAt   time.Time `json:"created_at"`
}