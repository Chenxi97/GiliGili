package dbops

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	dbConn *gorm.DB
	err    error
)

type VideoDel struct {
	ID string `gorm:"type:varchar(64);not_null"`
}

func init() {
	dbConn, err = gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/video_server?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	dbConn.SingularTable(true)
	dbConn.AutoMigrate(&VideoDel{})
}
