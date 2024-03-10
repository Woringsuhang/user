package grpcs

import (
	"flag"
	"fmt"
	"github.com/Woringsuhang/user/global"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"net"
)

func GetHostIp() string {
	addrList, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("get current host ip err: ", err)
		return ""
	}
	var ip string
	for _, address := range addrList {
		if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ip = ipNet.IP.String()
				break
			}
		}
	}
	return ip
}
func Registration(register func(g *grpc.Server), cert, key string) error {
	flag.Parse()
	//grpc中间件
	//creds, err := credentials.NewServerTLSFromFile(cert, key)
	//if err != nil {
	//	log.Fatalf("failed to create credentials: %v", err)
	//}
	//grpc.Creds(creds)
	g := grpc.NewServer()
	//ip := GetHostIp()
	listen, err := net.Listen(global.ConfigAll.Grpc.Agreement, fmt.Sprintf("%v:%s", "0.0.0.0", global.ConfigAll.Grpc.Port))
	if err != nil {
		zap.S().Panic(err)
	}
	reflection.Register(g)
	register(g)
	zap.S().Info("started...", listen.Addr())
	grpc_health_v1.RegisterHealthServer(g, health.NewServer())
	err = g.Serve(listen)
	if err != nil {
		return err
	} else {
		return nil
	}
}
