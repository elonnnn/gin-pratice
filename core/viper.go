package core

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// 初始化一个viper配置
func Viper() *viper.Viper {
	v := viper.New()
	//制定配置文件的路径
	v.SetConfigFile("config.yaml")
	v.SetConfigType("yaml")
	// 读取配置信息
	err := v.ReadInConfig()
	if err != nil {
		// 读取配置信息失败
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	//监听修改
	v.WatchConfig()
	//为配置修改增加一个回调函数
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了...")
	})
	return v
}
