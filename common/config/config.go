package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path"
)

func init() {
	InitConing()
}

func InitConing() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(path.Join(workDir, "config"))

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Print("获取配置文件错误")
		panic(err)
	} else {
		fmt.Print("配置初始化成功")
	}
}
