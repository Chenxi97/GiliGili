package main

import (
	// "html/template"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/statics", "./templates")
	//**代表文件夹，*代表文件
	//r.LoadHTMLGlob("../templates/*")
	r.LoadHTMLFiles("./templates/home.html", "./templates/userhome.html")

	r.GET("/", homeHandler)
	r.POST("/", homeHandler)
	r.GET("/userhome", userHomeHandler)
	r.POST("/userhome", userHomeHandler)
	r.POST("/api", apiHandler)
	//解决跨域问题
	r.POST("/upload/:vid-id", proxyHandler)
	r.GET("/videos/:vid-id", proxyHandler)

	r.Run(":8080")
}
