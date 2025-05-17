package nacos

import (
	"stuoj-api/infrastructure/client"
	"stuoj-api/interfaces/rpc/interceptors"
	"user-service/pkg/config"
)

func LoadConfig() error {
	var err error

	if err = NacosClient.GetConfig(&config.Conf.UserService); err != nil {
		return err
	}

	client.Token = config.Conf.UserService.Grpc.Token
	interceptors.Token = config.Conf.UserService.Grpc.Token

	return nil
}
