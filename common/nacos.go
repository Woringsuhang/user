package common

import (
	"encoding/json"
	"fmt"
	"github.com/Woringsuhang/user/global"
	"github.com/Woringsuhang/user/model"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"log"
)

func Nacos(dataId, group string) {
	//create clientConfig
	clientConfig := constant.ClientConfig{
		NamespaceId:         global.NacosConfig.NamespaceId, //we can create multiple clients with different namespaceId to support multiple namespace.When namespace is public, fill in the blank string here.
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}

	// At least one ServerConfig
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      global.NacosConfig.IpAddr,
			ContextPath: "/nacos",
			Port:        uint64(global.NacosConfig.Port),
			Scheme:      "http",
		},
	}
	// Creat Nacos Service Discovery Client
	namingClient, err := clients.CreateNamingClient(map[string]interface{}{
		"serverConfigs": []constant.ServerConfig{
			{
				IpAddr:      global.NacosConfig.IpAddr,
				Port:        uint64(global.NacosConfig.Port),
				ContextPath: "/nacos",
			},
		},
		"clientConfig": clientConfig,
	})

	if err != nil {
		log.Fatalf("Failed to create naming client: %v", err)
	}

	// Register Service Instance
	success, err := namingClient.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          "10.2.171.94",
		Port:        8080,
		ServiceName: "2108a",
		Weight:      1,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		Metadata:    map[string]string{},
	})
	if err != nil {
		log.Fatalf("Failed to register instance to nacos: %v", err)
	}

	log.Printf("Register instance success: %v", success)

	// Create config client for dynamic configuration
	configs, _ := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})
	fmt.Println(global.NacosConfig.DataId)
	fmt.Println(global.NacosConfig.Group)
	config, err := configs.GetConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group,
	})

	if err != nil {
		log.Printf("Error getting configuration")
		return
	}
	err = json.Unmarshal([]byte(config), &global.ConfigAll)
	fmt.Println(global.ConfigAll.Grpc.Port)
	fmt.Println(global.NacosConfig)
	//监听
	err = configs.ListenConfig(vo.ConfigParam{
		DataId: "configuration",
		Group:  "dev",
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("group:" + group + ",dataId:" + dataId + ",data:" + data)
			err = json.Unmarshal([]byte(data), &global.ConfigAll.Mysql)

			dns := model.Dns(global.ConfigAll.Mysql.Username, global.ConfigAll.Mysql.Password, global.ConfigAll.Mysql.Host,
				global.ConfigAll.Mysql.Port, global.ConfigAll.Mysql.Library)
			updateDbConnection(dns)
		},
	})

	if err != nil {
		panic(err)
	}
}
func updateDbConnection(config string) {
	// 关闭现有连接池（如果存在）
	Dbs, _ := model.DB.DB()
	if Dbs != nil {
		_ = Dbs.Close()
	}

	// 使用新的配置信息创建数据库连接
	var err error
	model.DB, err = gorm.Open(mysql.Open(config), &gorm.Config{})
	// 假设 config 是有效的数据库 DSN
	if err != nil {
		log.Fatalf("Failed to create database connection: %v", err)
	}

	// 可能需要对 db 进行额外配置，如设置连接池大小等

	fmt.Println("Database connection updated successfully.")
}
