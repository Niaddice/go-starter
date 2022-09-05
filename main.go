package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-starter/logger"
	routers "go-starter/router"
	"go-starter/src/auth"
	"go-starter/src/center"
	"go-starter/src/common"
	"go-starter/src/sso"
	"net/http"
	"runtime/debug"
)

func main() {
	fmt.Println("----> Start Server")
	//err := redis.InitClient()
	routers.Include(nc.Routers, center.Routers, sso.Routers)
	r := routers.Init()
	r.StaticFS("/upload", http.Dir("upload"))
	r.Use(Recover)
	//r.Use(middlewares.Cors())
	err := r.Run(":8092")
	if err != nil {
		logger.Log.Error("启动失败: %s ", err.Error())
	}

}

func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			//打印错误堆栈信息
			logger.Log.Info("panic: %v\n", r)
			debug.PrintStack()
			c.JSON(http.StatusOK, common.Error(errorToString(r)))
			logger.Log.Error(errorToString(r))
			c.JSON(http.StatusInternalServerError, gin.H{
				"Code": "-1",
				"Msg":  "内部错误",
			})
			//终止后续接口调用，不加的话recover到异常后，还会继续执行接口里后续代码
			c.Abort()
		}
	}()
	//加载完 defer recover，继续后续接口调用
	c.Next()
}

// recover错误，转string
func errorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return r.(string)
	}
}
