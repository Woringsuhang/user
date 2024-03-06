package common

import (
	"github.com/Woringsuhang/user/global"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func InitViper(path string) {
	//实例化一个viper方法
	v := viper.New()

	//viper.AutomaticEnv()

	//手动设置读取的文件路径

	v.SetConfigFile(path)

	//启用配置文件的动态监视,配置文件发生变化时自动重新加载配置
	v.WatchConfig()
	//读取配置文件
	err := v.ReadInConfig()
	if err != nil {
		zap.S().Panic(err)
		return
	}
	//把读取的配置文件信息拿出来
	err = v.Unmarshal(&global.NacosConfig)
	if err != nil {
		zap.S().Panic("解析yaml配置文件失败")
		return
	}

	//若配置文件发生了变化
	v.OnConfigChange(func(c fsnotify.Event) {
		//把读取的配置文件信息拿出来
		err = v.Unmarshal(&global.NacosConfig)
		if err != nil {
			zap.S().Panic("解析yaml配置文件失败")
			return
		}
		zap.S().Info("rpc配置发生变动")
	})
	zap.S().Info("viper初始化完成")
}
