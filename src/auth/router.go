package nc

import (
	"github.com/gin-gonic/gin"
	middleware2 "go-starter/src/middleware"
)

func Routers(e *gin.Engine) {

	auth := e.Group("/auth")
	{
		auth.POST("/token", middleware2.AuthHandler)
	}
}
