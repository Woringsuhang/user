package common

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
	"log"
	"net"
)

var ConsulCli *api.Client
var Srvid string

func ConsulClient(addr string) error {
	var err error

	ConsulCli, err = api.NewClient(&api.Config{
		Address: addr,
	})
	if err != nil {
		return errors.New("连接consul客户端失败！" + err.Error())
	}
	return nil
}
func GetIp() (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	for _, i := range interfaces {
		addrs, err := i.Addrs()
		if err != nil {
			continue
		}

		for _, addr := range addrs {
			ipNet, isVailIpNet := addr.(*net.IPNet)
			if isVailIpNet && !ipNet.IP.IsLoopback() {
				if ipNet.IP.To4() != nil {
					// 添加一些额外的检测逻辑，例如判断IP地址是否在本地网络范围内
					if ipNet.IP.IsGlobalUnicast() {
						// 添加详细的日志输出
						log.Printf("获取到的IP地址：%s，对应网络接口：%s\n", ipNet.IP.String(), i.Name)
						return ipNet.IP.String(), nil
					}
				}
			}
		}
	}

	return "", errors.New("Unable to find a valid global unicast IP address")
}
func AgentService(Address string, Port int) error {
	Srvid = uuid.New().String()
	ip, _ := GetIp()
	check := &api.AgentServiceCheck{
		Interval:                       "5s",
		Timeout:                        "5s",
		GRPC:                           fmt.Sprintf("%v:%v", ip, Port),
		DeregisterCriticalServiceAfter: "30s",
	}
	fmt.Println("wwwww", Port, ip)
	err := ConsulCli.Agent().ServiceRegister(&api.AgentServiceRegistration{
		ID:      Srvid,
		Name:    "user_srv",
		Tags:    []string{"GRPC"},
		Port:    Port,
		Address: ip,
		Check:   check,
	})
	if err != nil {
		return err
	}
	return nil
}

func GetConsul(serverName string) (ip string, port int, err error) {
	name, i, err := ConsulCli.Agent().AgentHealthServiceByName(serverName)
	fmt.Println(name)
	fmt.Println(i)
	if err != nil {
		return "", 0, err
	}
	var Ip string
	var Port int
	for _, val := range i {
		Ip = val.Service.Address
		Port = val.Service.Port
	}
	log.Println("端口：lianjie", Ip, Port)
	return Ip, Port, nil
}
