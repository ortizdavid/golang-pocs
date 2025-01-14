package entities

import (
	"time"
)

type Task struct {
	TaskId      	int `gorm:"primaryKey;autoIncrement"`
	UserId    		int `gorm:"column:user_id"`
	TaskName    	string `gorm:"column:task_name"`
	StartDate		string `gorm:"column:start_date"`
	EndDate			string  `gorm:"column:end_date"`
	Description 	string `gorm:"column:description"`
	Attachment 		string `gorm:"column:attachment"`
	UniqueId 		string `gorm:"column:unique_id"`
	CreatedAt  		time.Time `gorm:"column:created_at"`
	UpdatedAt  		time.Time `gorm:"column:updated_at"`
}

func (task Task) TableName() string {
	return "tasks"
}

type TaskData struct {
	TaskId			int
	UniqueId		string
	TaskName 		string
	StartDate		string
	EndDate			string
	Description		string
	Attachment		string 
	CreatedAt		string
	UpdatedAt		string
	UserId			int
	UserName		string
}

