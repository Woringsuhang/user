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

func Registration(register func(g *grpc.Server), cert, key string) error {
	flag.Parse()
	//grpc中间件
	//creds, err := credentials.NewServerTLSFromFile(cert, key)
	//if err != nil {
	//	log.Fatalf("failed to create credentials: %v", err)
	//}
	//grpc.Creds(creds)
	g := grpc.NewServer()
	listen, err := net.Listen(global.ConfigAll.Grpc.Agreement, fmt.Sprintf("0.0.0.0:%s", global.ConfigAll.Grpc.Port))
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
