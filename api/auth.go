package main

import (
	"log"

	"github.com/Chenxi97/GiliGili/api/defs"
	"github.com/Chenxi97/GiliGili/api/session"
	"github.com/gin-gonic/gin"
)

var HEADER_FIELD_SESSION = "X-Session-Id"
var HEADER_FIELD_UNAME = "X-User-Name"

func validateUserSession(c *gin.Context) bool {
	sid := c.GetHeader(HEADER_FIELD_SESSION)
	log.Println("validateUserSession:", sid)
	if len(sid) != 0 {
		if uname, ok := session.IsSessionExpired(sid); !ok {
			c.Request.Header.Add(HEADER_FIELD_UNAME, uname)
			return true
		}
	}
	return false
}

func ValidateUserSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		flag := validateUserSession(c)
		log.Print("Middle-ValidateSession:", flag)
	}
}

func ValidateUser(c *gin.Context) bool {
	uname := c.GetHeader(HEADER_FIELD_UNAME)
	log.Println("ValidateUser:", uname)
	if len(uname) == 0 {
		sendErrorResponse(c, &defs.ErrorNotAuthUser)
		return false
	}
	return true
}
