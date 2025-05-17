package interceptors

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"user-service/pkg/config"
)

func TokenAuthInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		conf := config.Conf.UserService.Grpc

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			log.Println("无法获取元数据，拒绝请求")
			return nil, status.Error(codes.Unauthenticated, "无法获取元数据")
		}

		token, ok := md["token"]
		if !ok || len(token) == 0 {
			log.Println("缺少token，拒绝请求")
			return nil, status.Error(codes.Unauthenticated, "缺少授权头")
		}

		if token[0] != conf.Token {
			log.Println("token 错误，拒绝请求")
			return nil, status.Errorf(codes.Unauthenticated, "token 错误，拒绝请求")
		}

		return handler(ctx, req)
	}
}
