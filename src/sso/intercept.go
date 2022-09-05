package sso

import (
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"go-starter/src/common"
	"net/http"
	"time"
)

var ca = cache.New(24*time.Hour, 25*time.Hour)

func HeaderCheck() func(c *gin.Context) {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		t := c.GetHeader("type")
		if t == "web" {
			// 获取header的值
			if token == "" {
				// header缺失
				c.JSON(http.StatusUnauthorized, common.Failed("登录超时"))
				c.Abort()
				return
			} else {
				_, exists := ca.Get(token)
				if !exists {
					ca.Delete(token)
					c.JSON(http.StatusUnauthorized, common.Failed("登录超时"))
					c.Abort()
					return
				}
			}
		}
		// 请求正常
		c.Next()
	}
}
