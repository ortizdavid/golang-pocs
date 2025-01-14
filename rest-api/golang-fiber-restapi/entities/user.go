package entities

import "time"

type User struct {
	UserId    int `gorm:"column:user_id;autoIncrement;primarykey"`
	UserName  string `gorm:"column:user_name;type:varchar(100)"`
	Password  string `gorm:"column:password;type:varchar(150)"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdateAt  time.Time `gorm:"column:updated_at"`
}

func TableName() string {
	return "users"
}