package center

import (
	"github.com/gin-gonic/gin"
)

func Routers(e *gin.Engine) {
	center := e.Group("/center")
	{
		msg := center.Group("/msg")
		{
			msg.POST("/listen", MsgListenHandler)
		}
	}
}
