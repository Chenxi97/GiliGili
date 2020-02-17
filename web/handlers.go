package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

type HomePage struct {
	Name string
}

type UserPage struct {
	Name string
}

func homeHandler(c *gin.Context) {
	cname, err1 := c.Cookie("username")
	sid, err2 := c.Cookie("session")

	//visitor
	if err1 != nil || err2 != nil {
		c.HTML(http.StatusOK, "home.html", &HomePage{
			Name: "visitor",
		})
		return
	}

	//user
	if len(cname) != 0 && len(sid) != 0 {
		c.Redirect(http.StatusFound, "/userhome")
		return
	}
}

func userHomeHandler(c *gin.Context) {
	//通过cookie进入
	cname, err1 := c.Cookie("username")
	_, err2 := c.Cookie("session")
	if err1 != nil || err2 != nil {
		c.Redirect(http.StatusFound, "/")
		return
	}
	//通过表单进入
	fname := c.PostForm("username")
	var p *UserPage
	if len(cname) != 0 {
		p = &UserPage{Name: cname}
	} else if len(fname) != 0 {
		p = &UserPage{Name: fname}
	}

	c.HTML(http.StatusOK, "userhome.html", p)
}

func proxyHandler(c *gin.Context) {
	u, _ := url.Parse("http://127.0.0.1:9000/")
	proxy := httputil.NewSingleHostReverseProxy(u)
	proxy.ServeHTTP(c.Writer, c.Request)
}

func apiHandler(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusBadRequest, ErrorRequestNotRecognized)
		return
	}

	apibody := &ApiBody{}
	if err := c.ShouldBind(apibody); err != nil {
		c.JSON(http.StatusBadRequest, ErrorRequestBodyParseFailed)
	}
	request(apibody, c.Writer, c.Request)
}
