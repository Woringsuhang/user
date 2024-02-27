package common

import (
	"encoding/json"
	"github.com/Woringsuhang/user/global"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"go.uber.org/zap"

	"log"
)

func Consul() {
	//create clientConfig
	clientConfig := constant.ClientConfig{
		NamespaceId:         "2bdf0290-9626-41e8-821f-00c954bc107e", //we can create multiple clients with different namespaceId to support multiple namespace.When namespace is public, fill in the blank string here.
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}

	// At least one ServerConfig
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      "127.0.0.1",
			ContextPath: "/nacos",
			Port:        8848,
			Scheme:      "http",
		},
	}

	// Create config client for dynamic configuration
	configs, _ := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})

	config, err := configs.GetConfig(vo.ConfigParam{
		DataId: "configuration",
		Group:  "dev",
	})

	if err != nil {
		log.Printf("Error getting configuration")
		return
	}

	err = json.Unmarshal([]byte(config), &global.ConfigAll)
	zap.S().Info(global.ConfigAll)
	if err != nil {
		panic(err)
	}
}
