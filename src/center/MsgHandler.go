package center

import (
	"github.com/gin-gonic/gin"
	"go-starter/src/common"
	"net/http"
)

func MsgListenHandler(c *gin.Context) {
	var msg Msg
	err := c.ShouldBind(&msg)
	if err != nil {
		c.JSON(http.StatusOK, common.Failed(common.RequestParamError))
	}
	//wechat.SendTextMsg(msg.Text, msg.UserId)
}
