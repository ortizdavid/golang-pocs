package models

import (
	"github.com/ortizdavid/golang-fiber-restapi/config"
	"github.com/ortizdavid/golang-fiber-restapi/entities"
	"gorm.io/gorm"
)

type TaskModel struct {
}

func (TaskModel) Create(task entities.Task) *gorm.DB {
	db, _ := config.ConnectDB()
	return db.Create(&task)
}

func (TaskModel) FindAll() []entities.Task {
	db, _ := config.ConnectDB()
	Tasks := []entities.Task{}
	db.Find(&Tasks)
	return Tasks
}

func (TaskModel) Update(task entities.Task) *gorm.DB {
	db, _ := config.ConnectDB()
	return db.Save(&task)
}

func (TaskModel) FindById(id int) entities.Task {
	db, _ := config.ConnectDB()
	var task entities.Task
	db.First(&task, id)
	return task
}

func (TaskModel) Delete(id int) *gorm.DB {
	db, _ := config.ConnectDB()
	var task entities.Task
	return db.Delete(&task, id)
}

func (TaskModel) Search(param string) []entities.Task {
	db, _ := config.ConnectDB()
	tasks := []entities.Task{}
	db.Where("task_name like ?", "%"+param+"%").Find(&tasks)
	return tasks
}

func (TaskModel) Exists(id int) bool {
	db, _ := config.ConnectDB()
	var task entities.Task
	result := db.First(&task, id)
	if result.RowsAffected == 0 {
		return false
	}
	return true
}
