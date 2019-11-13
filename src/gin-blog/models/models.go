//创建基础数据模型，初始化数据库
package models

import (
	"../pkg/setting"
	"fmt"
	"github.com/jinzhu/gorm"

	_"github.com/jinzhu/gorm/dialects/sqlite"
	_"github.com/go-sql-driver/mysql"
	"log"
)
var db *gorm.DB

type Model struct {
	ID int `gorm:"primary_key" json:"id"`
	CreatedOn int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}


func Setup() {
	var (
		err error
	)

	db, err = gorm.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name))

	if err != nil {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
		return setting.DatabaseSetting.TablePrefix + defaultTableName;
	}

	db.SingularTable(true)//全局禁用表名复数
	//db.DB().SetMaxIdleConns(10)
	//db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer db.Close()
}

