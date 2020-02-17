package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func streamHandler(c *gin.Context) {
	vid := c.Param("vid-id")
	vl := VIDEO_DIR + vid

	video, err := os.Open(vl)
	if err != nil {
		log.Printf("Error when try to open file: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	defer video.Close()
	//b, _ := ioutil.ReadAll(video)
	c.Header("Content-Type", "video/mp4")
	http.ServeContent(c.Writer, c.Request, "", time.Now(), video)
}

func uploadHandler(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["video_file"]

	for _, file := range files {
		log.Println(file.Filename)
		dst := fmt.Sprintf("./videos/%s", file.Filename)
		// 上传文件到指定的目录
		c.SaveUploadedFile(file, dst)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%d files uploaded!", len(files)),
	})
}
