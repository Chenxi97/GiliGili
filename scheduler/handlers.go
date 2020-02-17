package main

import (
	"github.com/Chenxi97/GiliGili/scheduler/dbops"
	"github.com/gin-gonic/gin"
)

func vidDelRecHandler(c *gin.Context) {
	vid := c.Param("vid-id")

	if len(vid) == 0 {
		c.JSON(400, gin.H{
			"error": "video id should not be empty",
		})
		return
	}

	err := dbops.AddVideoDeletionRecord(vid)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Internal server error",
		})
		return
	}
	c.JSON(200, gin.H{
		"success": "video deleted",
	})

}
