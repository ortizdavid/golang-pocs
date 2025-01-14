package models

import (
	"github.com/ortizdavid/go-rest-concepts/entities"
	"gorm.io/gorm"
)

type UserModel struct {
	db *gorm.DB
}

func NewUserModel(db *gorm.DB) *UserModel {
	return &UserModel{db: db}
}

func (model UserModel) Create(user entities.User) (*gorm.DB, error) {
	result := model.db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return result, nil
}

func (model UserModel) FindAll() ([]entities.User, error) {
	var users []entities.User
	result := model.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (model UserModel) Update(user entities.User) (*gorm.DB, error) {
	result := model.db.Save(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return result, nil
}

func (model UserModel) Delete(user entities.User) (*gorm.DB, error) {
	result := model.db.Delete(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return result, nil
}

func (model UserModel) FindById(id int) (entities.User, error) {
	var user entities.User
	result := model.db.First(&user, id)
	if result.Error != nil {
		return entities.User{}, result.Error
	}
	return user, nil
}

func (model UserModel) FindByUniqueId(uniqueId string) (entities.User, error) {
	var user entities.User
	result := model.db.First(&user, "unique_id=?", uniqueId)
	if result.Error != nil {
		return entities.User{}, result.Error
	}
	return user, nil
}

func (model UserModel) FindByUserName(userName string) (entities.User, error) {
	var user entities.User
	result := model.db.First(&user, "user_name=?", userName)
	if result.Error != nil {
		return entities.User{}, result.Error
	}
	return user, nil
}


func (model UserModel) Search(param interface{}) ([]entities.UserData, error) {
	var users []entities.UserData
	result := model.db.Raw("SELECT * FROM view_user_data WHERE user_name=?", param).Scan(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (model UserModel) Count() (int64) {
	var count int64
	result := model.db.Table("users").Count(&count)
	if result.Error != nil {
		return 0
	}
	return count
}

func (model UserModel) GetDataById(id int) (entities.UserData, error) {
	var userData entities.UserData
	result := model.db.Raw("SELECT * FROM view_user_data WHERE user_id=?", id).Scan(&userData)
	if result.Error != nil {
		return entities.UserData{}, result.Error
	}
	return userData, nil
}


func (model UserModel) FindAllData() ([]entities.UserData, error) {
	var users []entities.UserData
	result := model.db.Raw("SELECT * FROM view_user_data").Scan(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (model UserModel) FindAllDataLimit(start int, end int) ([]entities.UserData, error) {
	var users []entities.UserData
	result := model.db.Raw("SELECT * FROM view_user_data LIMIT ?, ?", start, end).Scan(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}


func (model UserModel) ExistsRecord(fieldName string, value any) (bool, error) {
	var user entities.User
	result := model.db.Where(fieldName+" = ?", value).First(&user)
	if result.Error != nil {
		return false, result.Error
	}
	return user.UserId != 0, nil
}
