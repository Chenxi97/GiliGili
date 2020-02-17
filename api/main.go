package main

import (
	"github.com/Chenxi97/GiliGili/api/session"
	"github.com/gin-gonic/gin"
)

func prepare() {
	session.LoadSessionsFromDB()
}

func main() {
	prepare()
	r := gin.Default()
	r.Use(ValidateUserSession())

	r.POST("/user", CreateUser)
	r.POST("/user/:username", Login)
	r.GET("/user/:username", GetUserInfo)
	r.POST("/user/:username/videos", AddNewVideo)
	r.GET("/user/:username/videos", ListAllVideos)
	r.DELETE("/user/:username/videos/:vid-id", DeleteVideo)
	r.POST("/videos/:vid-id/comments", PostComment)
	r.GET("/videos/:vid-id/comments", ShowComments)

	r.Run(":8000")
}
