package grpcs

import (
	"flag"
	"fmt"
	"github.com/Woringsuhang/user/global"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"net"
)

func Registration(register func(g *grpc.Server)) error {
	flag.Parse()
	g := grpc.NewServer()

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
