package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"keep-bill/pkg/setting"
	"time"
)

var db *gorm.DB

type Model struct {
	// 默认model结构体
	ID         int       `gorm:"primary_key" json:"id"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

func init() {
	var (
		err                                  error
		dbType, dbName, user, password, host string
	)

	dbType = setting.Type
	dbName = setting.Name
	user = setting.User
	password = setting.Password
	host = setting.Host

	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	if err != nil {
		panic("连接数据库失败")
	}

	db.SingularTable(true) // 禁用复数表
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

}
