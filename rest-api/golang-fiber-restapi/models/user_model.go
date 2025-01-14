package models

import (
	"github.com/ortizdavid/golang-fiber-restapi/config"
	"github.com/ortizdavid/golang-fiber-restapi/entities"
	"gorm.io/gorm"
)

type UserModel struct {
}

func (UserModel) Create(user entities.User) *gorm.DB {
	db, _ := config.ConnectDB()
	return db.Create(&user)
}

func (UserModel) FindAll() []entities.User {
	db, _ := config.ConnectDB()
	Users := []entities.User{}
	db.Find(&Users)
	return Users
}

func (UserModel) Update(user entities.User) *gorm.DB {
	db, _ := config.ConnectDB()
	return db.Save(&user)
}

func (UserModel) FindById(id int) entities.User {
	db, _ := config.ConnectDB()
	var user entities.User
	db.First(&user, id)
	return user
}

func (UserModel) Delete(id int) *gorm.DB {
	db, _ := config.ConnectDB()
	var user entities.User
	return db.Delete(&user, id)
}

func (UserModel) Search(param string) []entities.User {
	db, _ := config.ConnectDB()
	users := []entities.User{}
	db.Where("user_name like ?", "%"+param+"%").Find(&users)
	return users
}

func (UserModel) Exists(id int) bool {
	db, _ := config.ConnectDB()
	var user entities.User
	result := db.First(&user, id)
	if result.RowsAffected == 0 {
		return false
	}
	return true
}

func (UserModel) FindByCredentials(username string, password string) entities.User {
	db, _ := config.ConnectDB()
	var user entities.User
	db.Where("user_name = ? AND password = ?", username, password).First(&user)
	return user
}