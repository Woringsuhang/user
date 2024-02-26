package common

import "go.uber.org/zap"

func InitZap() {
	//创建一个日志记录器
	development, err := zap.NewDevelopment()
	if err != nil {
		panic(err.Error())
	}
	//数据同步
	defer development.Sync()
	//动态切换日志记录器的配置
	zap.ReplaceGlobals(development)
}
