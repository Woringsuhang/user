package common

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"

	"log"
)

var ConsulCli *api.Client
var Srvid string

func ConsulClient() error {
	var err error

	ConsulCli, err = api.NewClient(&api.Config{
		Address: "10.2.171.94:8085",
	})
	if err != nil {
		return errors.New("连接consul客户端失败！" + err.Error())
	}
	return nil
}

func AgentService(Address string, Port int) error {
	Srvid = uuid.New().String()

	check := &api.AgentServiceCheck{
		Interval:                       "5s",
		Timeout:                        "5s",
		GRPC:                           fmt.Sprintf("%s:%d", Address, Port),
		DeregisterCriticalServiceAfter: "30s",
	}
	err := ConsulCli.Agent().ServiceRegister(&api.AgentServiceRegistration{
		ID:      Srvid,
		Name:    "user_srv",
		Tags:    []string{"GRPC"},
		Port:    Port,
		Address: Address,
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