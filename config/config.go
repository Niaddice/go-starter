package config

import (
	"fmt"
	"github.com/bluele/gcache"
	"github.com/spf13/viper"
	"go-starter/logger"
	"os"
	"time"
)

type Config struct {
	encodingAesKey string
	token          string
	orgcode        string
}

var gc = gcache.New(20).LFU().Build()

func GetConfig(key string) interface{} {
	//第一步先获取主配置文件中的active配置
	active, err := gc.Get("active")
	if err != nil {
		workDir, _ := os.Getwd()
		viper.AddConfigPath(workDir + "/resource")
		//设置读取的文件名
		viper.SetConfigName("conf")
		//设置要读取的文件类型
		viper.SetConfigType("yml")
		err := viper.ReadInConfig() // 搜索并读取配置文件
		if err != nil {             // 处理错误
			logger.Log.Error(fmt.Errorf("Fatal error config file: %s \n", err))
		}
		get := viper.Get("active")
		gc.Set("active", get)
		gc.SetWithExpire("active", get, time.Minute*60)
		active = get
	}
	//第二部获取对应环境配置中的具体配置信息
	value, err2 := gc.Get(key)
	if err2 != nil {
		workDir, _ := os.Getwd()
		viper.AddConfigPath(workDir + "/resource")
		//设置读取的文件名
		viper.SetConfigName("conf-" + active.(string))
		//设置要读取的文件类型
		viper.SetConfigType("yml")
		err := viper.ReadInConfig() // 搜索并读取配置文件
		if err != nil {             // 处理错误
			logger.Log.Error(fmt.Errorf("Fatal error config file: %s \n", err))
		}
		get := viper.Get(key)
		gc.Set(key, get)
		gc.SetWithExpire(key, get, time.Minute*60)
		return get
	} else {
		return value
	}
}
