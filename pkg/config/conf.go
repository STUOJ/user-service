package config

import (
	"stuoj-common/pkg/config"
	"stuoj-common/pkg/utils"
)

var (
	Conf *Config = &Config{}
)

type Config struct {
	Nacos       config.NacosConf `yaml:"nacos" json:"nacos"`
	UserService UserServiceConf  `yaml:"user-service" json:"user_service"`
}

// Config 初始化
func InitConfig() error {
	v, err := utils.IsFileExists("config.yaml")
	if err != nil {
		return err
	}
	if !v {
		Conf.Default()
		err = utils.WriteYaml(&Conf, "config.yaml")
		if err != nil {
			return err
		}
	}
	err = utils.ReadYaml(&Conf, "config.yaml")
	if err != nil {
		return err
	}
	utils.EmailHost = Conf.UserService.Email.Host
	utils.EmailPort = Conf.UserService.Email.Port
	utils.FromEmail = Conf.UserService.Email.Email
	utils.FromEmailSmtpPwd = Conf.UserService.Email.SmtpPwd
	return nil
}

func (c *Config) Default() {
	c.Nacos.Default()
}
