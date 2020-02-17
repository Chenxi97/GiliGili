package main

import (
	"github.com/gin-gonic/gin"
)

var CL *ConnLimiter

func init() {
	CL = NewConnLimiter(2)
}

func main() {
	r := gin.Default()
	//流量控制
	r.Use(ConnLimit())
	r.GET("/videos/:vid-id", streamHandler)
	r.POST("/upload/:vid-id", uploadHandler)
	r.Run(":9000")
}
