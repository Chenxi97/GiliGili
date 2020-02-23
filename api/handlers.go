package main

import (
	"log"
	"net/http"

	"github.com/Chenxi97/GiliGili/api/dbops"
	"github.com/Chenxi97/GiliGili/api/defs"
	"github.com/Chenxi97/GiliGili/api/session"
	"github.com/Chenxi97/GiliGili/api/utils"
	"github.com/gin-gonic/gin"
)

func sendErrorResponse(c *gin.Context, r *defs.ErrResponse) {
	c.JSON(r.HttpSC, r.Error)
}

func CreateUser(c *gin.Context) {
	//读取消息体获取用户名和密码，并判断是否合法
	user := defs.User{}
	if err := c.BindJSON(&user); err == nil && user.LoginName != "" && user.Pwd != "" {
		log.Printf("user info:%#v\n", user)
	} else {
		log.Println(err.Error())
		sendErrorResponse(c, &defs.ErrorRequestBodyParseFailed)
		return
	}
	//写入数据库
	if _, err := dbops.AddUser(user.LoginName, user.Pwd); err != nil {
		log.Println(err.Error())
		sendErrorResponse(c, &defs.ErrorDBError)
		return
	}
	//生成session
	id := session.GenerateNewSessionId(user.LoginName)
	su := defs.SignedUp{Success: true, SessionID: id}

	c.JSON(http.StatusCreated, &su)
}

func Login(c *gin.Context) {
	//得到消息体
	user := defs.User{}
	if err := c.BindJSON(&user); err == nil && len(user.LoginName) != 0 && len(user.Pwd) != 0 {
		log.Printf("user info:%#v\n", user)
	} else {
		log.Println(err.Error())
		sendErrorResponse(c, &defs.ErrorRequestBodyParseFailed)
		return
	}
	//验证姓名是否相同
	uname := c.Param("username")
	if uname != user.LoginName {
		sendErrorResponse(c, &defs.ErrorNotAuthUser)
		return
	}
	//验证密码是否相同
	dbuser, err := dbops.GetUser(user.LoginName)
	if err != nil || len(dbuser.Pwd) == 0 || dbuser.Pwd != user.Pwd {
		sendErrorResponse(c, &defs.ErrorNotAuthUser)
		return
	}
	//添加session
	sid := session.GenerateNewSessionId(user.LoginName)
	si := &defs.SignedIn{Success: true, SessionID: sid}
	c.JSON(http.StatusOK, si)
}

func GetUserInfo(c *gin.Context) {
	if !ValidateUser(c) {
		log.Printf("Unauthorized user\n")
		return
	}

	uname := c.Param("username")
	u, err := dbops.GetUser(uname)
	if err != nil {
		log.Printf("Error in GetUserInfo: %s", err)
		sendErrorResponse(c, &defs.ErrorDBError)
		return
	}
	u.Pwd = ""
	c.JSON(http.StatusOK, u)
}

func AddNewVideo(c *gin.Context) {
	if !ValidateUser(c) {
		log.Printf("Unathorized user \n")
		return
	}

	//得到消息体
	nvbody := defs.VideoInfo{}
	if err := c.BindJSON(&nvbody); err == nil {
		log.Printf("user info:%#v\n", nvbody)
	} else {
		log.Println(err.Error())
		sendErrorResponse(c, &defs.ErrorRequestBodyParseFailed)
		return
	}

	//写入数据库
	vi, err := dbops.AddNewVideo(nvbody.AuthorID, nvbody.Name)
	log.Printf("Author id : %d, name: %s \n", nvbody.AuthorID, nvbody.Name)
	if err != nil {
		log.Printf("Error in AddNewVideo: %s", err)
		sendErrorResponse(c, &defs.ErrorDBError)
		return
	}
	c.JSON(http.StatusCreated, vi)
}

func ListAllVideos(c *gin.Context) {
	if !ValidateUser(c) {
		return
	}

	uname := c.Param("username")
	vs, err := dbops.ListVideoInfo(uname, 0, utils.GetCurrentTimestampSec())
	if err != nil {
		log.Printf("Error in ListAllvideos: %s", err)
		sendErrorResponse(c, &defs.ErrorDBError)
		return
	}

	vsi := &defs.VideosInfo{Videos: vs}
	c.JSON(http.StatusCreated, vsi)
}

func DeleteVideo(c *gin.Context) {
	if !ValidateUser(c) {
		return
	}

	vid := c.Param("vid-id")
	log.Print("DeleteVideo:", vid)
	//删除数据库中文件
	err := dbops.DeleteVideoInfo(vid)
	if err != nil {
		log.Printf("Error in DeletVideo: %s", err)
		sendErrorResponse(c, &defs.ErrorDBError)
		return
	}

	//由scheduler删除本地文件
	go utils.SendDeleteVideoRequest(vid)
	c.JSON(http.StatusNoContent, gin.H{})
}

func PostComment(c *gin.Context) {
	if !ValidateUser(c) {
		return
	}

	//读消息体
	cbody := &defs.CommentForList{}
	if err := c.BindJSON(&cbody); err == nil {
		log.Printf("user info:%#v\n", cbody)
	} else {
		log.Println(err.Error())
		sendErrorResponse(c, &defs.ErrorRequestBodyParseFailed)
		return
	}

	//写入数据库
	vid := c.Param("vid-id")
	if err := dbops.AddNewComments(vid, cbody.AuthorID, cbody.Content); err != nil {
		log.Printf("Error in PostComment: %s", err)
		sendErrorResponse(c, &defs.ErrorDBError)
		return
	}
	c.JSON(http.StatusCreated, gin.H{})
}

func ShowComments(c *gin.Context) {
	if !ValidateUser(c) {
		return
	}

	vid := c.Param("vid-id")
	cm, err := dbops.ListComments(vid, 0, utils.GetCurrentTimestampSec())
	if err != nil {
		log.Printf("Error in ShowComments: %s", err)
		sendErrorResponse(c, &defs.ErrorDBError)
		return
	}

	cms := &defs.Comments{Comments: cm}
	log.Print("ShowComments:", cms)
	c.JSON(http.StatusOK, cms)
}
