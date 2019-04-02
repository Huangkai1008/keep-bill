package db

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	/**
	用户
	*/
	Model
	Username string `json:"username"`
	Password string `json:"password"`
}

func CreateUser(user User) bool {
	// 创建用户
	db.Create(&user)
	return true
}

func ExistUser(params map[string]interface{}) bool {
	// 是否存在用户
	var count int
	username, existUsername := params["username"]

	if existUsername {
		db.Model(&User{}).Where("username = ?", username).Count(&count)
	}
	if count > 0 {
		return true
	}

	return false
}

func (user *User) BeforeCreate(scope *gorm.Scope) error {
	_ = scope.SetColumn("CreateTime", time.Now())
	return nil
}

func (user *User) BeforeUpdate(scope *gorm.Scope) error {
	_ = scope.SetColumn("UpdateTime", time.Now())
	return nil
}
