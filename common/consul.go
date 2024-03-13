package common

import (
	"errors"
	"fmt"
	"github.com/Woringsuhang/user/global"
	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"

	"net"
)

func ConsulClient(addr string) (error, *api.Client) {
	var err error

	ConsulCli, err := api.NewClient(&api.Config{
		Address: addr,
	})
	if err != nil {
		return errors.New("连接consul客户端失败！" + err.Error()), nil
	}
	return nil, ConsulCli
}

//	func GetIp() (string, error) {
//		interfaces, err := net.Interfaces()
//		if err != nil {
//			return "", err
//		}
//
//		for _, i := range interfaces {
//			addrs, err := i.Addrs()
//			if err != nil {
//				continue
//			}
//
//			for _, addr := range addrs {
//				ipNet, isVailIpNet := addr.(*net.IPNet)
//				if isVailIpNet && !ipNet.IP.IsLoopback() {
//					if ipNet.IP.To4() != nil {
//						// 添加一些额外的检测逻辑，例如判断IP地址是否在本地网络范围内
//						if ipNet.IP.IsGlobalUnicast() {
//							// 添加详细的日志输出
//							log.Printf("获取到的IP地址：%s，对应网络接口：%s\n", ipNet.IP.String(), i.Name)
//							return ipNet.IP.String(), nil
//						}
//					}
//				}
//			}
//		}
//
//		return "", errors.New("Unable to find a valid global unicast IP address")
//	}
func GetIp() (ip []string) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ip
	}
	for _, addr := range addrs {
		ipNet, isVailIpNet := addr.(*net.IPNet)
		if isVailIpNet && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ip = append(ip, ipNet.IP.String())
			}
		}

	}
	return ip
}
func ConnectionConsul() (*api.Client, error) {
	return api.NewClient(&api.Config{Address: fmt.Sprintf("%v:%v", global.ConfigAll.Consuls.Host, global.ConfigAll.Consuls.Port)})
}
func AgentService(port int, serviceName string) error {
	ip := GetIp()
	conn, err := ConnectionConsul()
	if err != nil {
		return err
	}
	fmt.Println(ip[0], "wwww")
	return conn.Agent().ServiceRegister(&api.AgentServiceRegistration{
		ID:      uuid.New().String(),
		Name:    serviceName,
		Tags:    []string{"GRPC"},
		Port:    port,
		Address: ip[0],
		Check: &api.AgentServiceCheck{
			Interval:                       "5s",
			Timeout:                        "5s",
			GRPC:                           fmt.Sprintf("%v:%v", ip[0], port),
			DeregisterCriticalServiceAfter: "30s",
		},
	})

}

func GetConsul(serverName string) (string, error) {
	consul, err := ConnectionConsul()
	if err != nil {
		return "", err
	}
	name, i, err := consul.Agent().AgentHealthServiceByName(serverName)
	fmt.Println(name)
	fmt.Println(i)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v:%v", i[0].Service.Address, i[0].Service.Port), nil
}
