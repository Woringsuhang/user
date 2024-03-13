package grpcs

import (
	"flag"
	"fmt"
	"github.com/Woringsuhang/user/common"
	"github.com/Woringsuhang/user/global"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"net"
	"strconv"
)

func GetFreePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer l.Close()

	return l.Addr().(*net.TCPAddr).Port, nil
}
func Registration(register func(g *grpc.Server), cert, key string) error {
	flag.Parse()
	//port, _ := common.GetFreePort()

	//grpc中间件
	//creds, err := credentials.NewServerTLSFromFile(cert, key)
	//if err != nil {
	//	log.Fatalf("failed to create credentials: %v", err)
	//}
	//grpc.Creds(creds)
	g := grpc.NewServer()
	//ip := GetHostIp()
	//port, err := GetFreePort()
	//if err != nil {
	//	panic(err)
	//}
	listen, err := net.Listen(global.ConfigAll.Grpc.Agreement, fmt.Sprintf("%v:%v", "0.0.0.0", global.ConfigAll.Grpc.Port))
	if err != nil {
		zap.S().Panic(err)
	}
	port, _ := strconv.Atoi(global.ConfigAll.Grpc.Port)

	err = common.AgentService(port, "server")
	if err != nil {
		return err
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
