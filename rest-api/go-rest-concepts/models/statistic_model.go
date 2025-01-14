package models

import "gorm.io/gorm"

type StatisticCount struct {
	Users 				int64
	Tasks 				int64
}

func NewStatisticsCount(db *gorm.DB) StatisticCount {
	countUsers :=  NewUserModel(db).Count()
	countTasks := NewTaskModel(db).Count()

	return StatisticCount{
		Users:           countUsers,
		Tasks:           countTasks,
	}
}
