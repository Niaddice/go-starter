package sso

import (
	"github.com/gin-gonic/gin"
)

func Routers(e *gin.Engine) {

	auth := e.Group("/sso")
	{
		auth.POST("/login", LoginHandler)
	}
}
