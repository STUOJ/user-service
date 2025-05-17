package nacos

import "user-service/pkg/config"

func LoadConfig() error {
	var err error

	if err = NacosClient.GetConfig(&config.Conf.UserService); err != nil {
		return err
	}

	return nil
}
