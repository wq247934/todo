package common

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"todo/datamodels"
)

var DB *gorm.DB

func init() {
	DB = NewMysqlConn()
}

//获取MySQL连接
func GetMysqlConn() *gorm.DB {
	return NewMysqlConn()
}

//创建mysql 连接
func NewMysqlConn() *gorm.DB {
	dsn := "root:123456@tcp(localhost:3306)/todo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dsn)
	db.AutoMigrate(&datamodels.Todo{})
	if err != nil {
		panic(err)
	}
	return db
}
