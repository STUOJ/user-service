package nacos

import "user-service/pkg/config"

func LoadConfig() error {
	var err error

	if err = NacosClient.GetConfig(&config.Conf.Database, "user-service-database.yaml"); err != nil {
		return err
	}

	if err = NacosClient.GetConfig(&config.Conf.Grpc, "user-service-grpc.yaml"); err != nil {
		return err
	}

	return nil
}
