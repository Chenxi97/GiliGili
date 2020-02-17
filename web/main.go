package main

import (
	// "html/template"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/statics", "./templates")
	//**代表文件夹，*代表文件
	//router.LoadHTMLGlob("../templates/*")
	router.LoadHTMLFiles("./templates/home.html", "./templates/userhome.html")

	router.GET("/", homeHandler)
	router.POST("/", homeHandler)
	router.GET("/userhome", userHomeHandler)
	router.POST("/userhome", userHomeHandler)
	router.POST("/api", apiHandler)
	//解决跨域问题
	router.POST("/upload/:vid-id", proxyHandler)

	router.Run(":8080")
}
