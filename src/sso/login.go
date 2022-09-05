package sso

import (
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"go-starter/db"
	"go-starter/src/common"
	"go-starter/src/middleware"
	"go-starter/src/utils"
	"net/http"
	"strconv"
	"time"
)

func LoginHandler(c *gin.Context) {
	// 用户发送用户名和密码过来
	var user User
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusOK, common.FailedByReqParam())
		return
	}
	var u User
	db.Db.Table("user").Where("username = ? and deleted_at is null and status = ?", user.Username, 1).First(&u)
	if u.Id != "" {
		//salt := utils.RandSalt(4)
		crypted := utils.GenerateHashPassword(user.Password, u.Salt)
		if crypted == u.Password {
			tokenString, _ := middleware.GenToken(user.Username + u.Salt + strconv.FormatInt(time.Now().Unix(), 10))
			db.Db.Table("user").Where("username = ? and deleted_at is null and status = ?", user.Username, 1).Update("token", tokenString)
			ca.Set(tokenString, u.Username, cache.DefaultExpiration)
			c.JSON(http.StatusOK, common.Ok(map[string]string{"token": tokenString, "name": u.DisplayName}))
			return
		} else {
			c.JSON(http.StatusOK, common.Failed("密码错误"))
			return
		}
	}
	c.JSON(http.StatusOK, common.Failed("用户不存在"))
	return
}
