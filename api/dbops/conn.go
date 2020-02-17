package dbops

import (
	"github.com/Chenxi97/GiliGili/api/defs"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	dbConn *gorm.DB
	err    error
)

func init() {
	dbConn, err = gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/video_server?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	dbConn.SingularTable(true)
	dbConn.AutoMigrate(&defs.User{})
	dbConn.AutoMigrate(&defs.VideoInfo{})
	dbConn.AutoMigrate(&defs.Comment{})
	dbConn.AutoMigrate(&defs.Session{})
}
