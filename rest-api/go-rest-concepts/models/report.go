package models

import "gorm.io/gorm"


type TableReport struct {
	Title  	string
	Rows	interface{}
	Count   int
	db 		*gorm.DB
}

func NewTableReport(db *gorm.DB) TableReport {
	return TableReport{db: db}
}


func (report TableReport) GetAllUsers() TableReport {
	rows, _ := NewUserModel(report.db).FindAllData()
	count := len(rows)
	return TableReport {
		Title: "Users",
		Rows: rows,
		Count: count,
	}
}


func (report TableReport) GetAllTasks() TableReport {
	rows, _ := NewUserModel(report.db).FindAllData()
	count := len(rows)
	return TableReport {
		Title: "Tasks",
		Rows: rows,
		Count: count,
	}
}
