package client

import (
	"stuoj-api/infrastructure/client"
	"user-service/pkg/config"
)

func Init() {
	client.Token = config.Conf.UserService.Grpc.Token
}
