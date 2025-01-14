package models

import (
	"github.com/ortizdavid/go-rest-concepts/entities"
	"gorm.io/gorm"
)

type TaskModel struct {
	db *gorm.DB
}

func NewTaskModel(db *gorm.DB) *TaskModel {
	return &TaskModel{db: db}
}

func (model TaskModel) Create(task entities.Task) (*gorm.DB, error) {
	result := model.db.Create(&task)
	if result.Error != nil {
		return nil, result.Error
	}
	return result, nil
}

func (model TaskModel) CreateBatch(tasks []entities.Task) (*gorm.DB, error) {
	tx := model.db.Begin()
	result := model.db.Create(&tasks)
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}
	tx.Commit()
	return result, nil
}

func (model TaskModel) FindAll() ([]entities.Task, error) {
	var tasks []entities.Task
	result := model.db.Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}

func (model TaskModel) Update(task entities.Task) (*gorm.DB, error) {
	result := model.db.Save(&task)
	if result.Error != nil {
		return nil, result.Error
	}
	return result, nil
}

func (model TaskModel) Delete(task entities.Task) (*gorm.DB, error) {
	result := model.db.Delete(&task)
	if result.Error != nil {
		return nil, result.Error
	}
	return result, nil
}

func (model TaskModel) FindById(id int) (entities.Task, error) {
	var task entities.Task
	result := model.db.First(&task, id)
	if result.Error != nil {
		return entities.Task{}, result.Error
	}
	return task, nil
}

func (model TaskModel) FindByUniqueId(uniqueId string) (entities.Task, error) {
	var task entities.Task
	result := model.db.First(&task, "unique_id=?", uniqueId)
	if result.Error != nil {
		return entities.Task{}, result.Error
	}
	return task, nil
}

func (model TaskModel) Search(param string) ([]entities.TaskData, error) {
	var tasks []entities.TaskData
	result := model.db.Raw("SELECT * FROM view_task_data WHERE task_name=?", param).Scan(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}

func (model TaskModel) GetDataById(id int) (entities.TaskData, error) {
	var taskData entities.TaskData
	result := model.db.Raw("SELECT * FROM view_task_data WHERE task_id=?", id).Scan(&taskData)
	if result.Error != nil {
		return entities.TaskData{}, result.Error
	}
	return taskData, nil
}

func (model TaskModel) GetDataByUniqueId(uniqueId string) (entities.TaskData, error) {
	var taskData entities.TaskData
	result := model.db.Raw("SELECT * FROM view_task_data WHERE unique_id=?", uniqueId).Scan(&taskData)
	if result.Error != nil {
		return entities.TaskData{}, result.Error
	}
	return taskData, nil
}

func (model TaskModel) FindAllData() ([]entities.TaskData, error) {
	var tasks []entities.TaskData
	result := model.db.Raw("SELECT * FROM view_task_data").Scan(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}

func (model TaskModel) FindAllDataLimit(start int, end int) ([]entities.TaskData, error) {
	var tasks []entities.TaskData
	result := model.db.Raw("SELECT * FROM view_task_data LIMIT ?, ?", start, end).Scan(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}

func (model TaskModel) FindAllDataByUserIdLimit(userId int, start int, end int) ([]entities.TaskData, error) {
	var tasks []entities.TaskData
	result := model.db.Raw("SELECT * FROM view_task_data WHERE user_id=? LIMIT ?, ?", userId, start, end).Scan(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}

func (model TaskModel) Count() int64 {
	var count int64
	result := model.db.Table("tasks").Count(&count)
	if result.Error != nil {
		return 0
	}
	return count
}

func (model TaskModel) ExistsRecord(fieldName string, value any) (bool, error) {
	var task entities.Task
	result := model.db.Where(fieldName+" = ?", value).First(&task)
	if result.Error != nil {
		return false, result.Error
	}
	return task.TaskId != 0, nil
}
