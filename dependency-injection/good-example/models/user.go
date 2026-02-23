package models

import "time"

type User struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    Name      string    `gorm:"size:100;not null" json:"name"`
    Email     string    `gorm:"uniqueIndex;not null" json:"email"`
    Password  string    `json:"-"` 
    CreatedAt time.Time `json:"created_at"`
}