package entities

import (
	"time"
)

type Task struct {
	TaskId      int `gorm:"primaryKey;autoIncrement"`
	TaskName    string `gorm:"column:task_name;type:varchar(100)"`
	StartDate	time.Time `gorm:"column:start_date"`
	EndDate		time.Time  `gorm:"column:end_date"`
	Description string `gorm:"column:description;type:varchar(300)"`
	Status  	string `gorm:"column:status;type:varchar(20)"`
}

func (task Task) TableName() string {
	return "tasks"
}
