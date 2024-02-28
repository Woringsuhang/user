package grpcs

import (
	"flag"
	"fmt"
	"github.com/Woringsuhang/user/global"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"google.golang.org/grpc/reflection"
	"log"

	"net"
)

func Registration(register func(g *grpc.Server), cert, key string) error {
	flag.Parse()
	//grpc中间件
	creds, err := credentials.NewServerTLSFromFile(cert, key)
	if err != nil {
		log.Fatalf("failed to create credentials: %v", err)
	}
	g := grpc.NewServer(grpc.Creds(creds))

	listen, err := net.Listen(global.ConfigAll.Grpc.Agreement, fmt.Sprintf(":%s", global.ConfigAll.Grpc.Port))
	if err != nil {
		zap.S().Panic(err)
	}
	reflection.Register(g)
	register(g)
	zap.S().Info("started...", listen.Addr())

	err = g.Serve(listen)
	if err != nil {
		return err
	} else {
		return nil
	}
}
