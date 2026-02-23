package infra

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)


func NewDatabase(dsn string) *gorm.DB {
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database: " + err.Error())
    }
    return db
}

