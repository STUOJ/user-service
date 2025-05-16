package config

import (
	"common/pkg/config"
	"common/pkg/utils"
)

var (
	Conf *Config = &Config{}
)

type Config struct {
	Database config.DatabaseConf `yaml:"database" json:"database"`
	Server   config.ServerConf   `yaml:"server" json:"server"`
	Email    EmailConf           `yaml:"email" json:"email"`
	Token    TokenConf           `yaml:"token" json:"token"`
	Nacos    config.NacosConf    `yaml:"nacos" json:"nacos"`
	Grpc     config.GrpcConf     `yaml:"grpc" json:"grpc"`
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
	utils.Expire = Conf.Token.Expire
	utils.Secret = Conf.Token.Secret
	utils.Refresh = Conf.Token.Refresh
	utils.EmailHost = Conf.Email.Host
	utils.EmailPort = Conf.Email.Port
	utils.FromEmail = Conf.Email.Email
	utils.FromEmailSmtpPwd = Conf.Email.SmtpPwd
	return nil
}

func (c *Config) Default() {
	c.Database.Default()
	c.Server.Default()
	c.Token.Default()
	c.Nacos.Default()
}
