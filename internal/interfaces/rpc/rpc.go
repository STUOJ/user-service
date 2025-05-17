package rpc

import (
	"google.golang.org/grpc"
	"net"
	"stuoj-api/api/pb"
	"user-service/internal/interfaces/rpc/handler"
	"user-service/internal/interfaces/rpc/interceptors"
	"user-service/pkg/config"
)

func Init() error {
	conf := config.Conf.UserService.Grpc

	listen, err := net.Listen("tcp", ":"+conf.Port)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interceptors.TokenAuthInterceptor(),
		),
	)

	// 注册服务
	pb.RegisterUserServiceServer(grpcServer, &handler.UserServer{})

	err = grpcServer.Serve(listen)
	if err != nil {
		return err
	}

	return nil
}
