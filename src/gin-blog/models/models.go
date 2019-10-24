//创建基础数据模型，初始化数据库
package models
import (
	"../pkg/setting"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
)
type Model struct {
	//gorm.Model  自带模型
	ID int `gorm:"primary_key" json:"id"`
	CreatedOn int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

var db *gorm.DB
func init(){
	var (
		err error
		dbType, dbName, user, password, host, tablePrefix string
	)
	sec,err := setting.Cfg.GetSection("database")
	if err !=nil{
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}
	dbType =sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()
	// db, err := gorm.Open("mysql", "user:password@(localhost)/dbname?charset=utf8&parseTime=True&loc=Local")
	db, err := gorm.Open(dbType,fmt.Sprintf("%s:%s@/tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user,
		password,
		host,
		dbName))
	if err!= nil {
		log.Println(err)
	}
	gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
		return tablePrefix + defaultTableName;
	}//修改默认表名处理程序

	db.SingularTable(true)  //使用单元格表
	//db.DB().SetMaxIdleConns(10)
	//db.DB().SetMaxOpenConns(100)
}

func CloseDB(){
	defer db.Close()
}