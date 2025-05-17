package config

import "stuoj-common/pkg/config"

type UserServiceConf struct {
	Grpc     config.GrpcConf     `yaml:"grpc" json:"grpc"`
	Server   config.ServerConf   `yaml:"server" json:"server"`
	Database config.DatabaseConf `yaml:"database" json:"database"`
	Email    EmailConf           `yaml:"email" json:"email"`
	Token    TokenConf           `yaml:"token" json:"token"`
}

func (c *UserServiceConf) Default() {
	c.Grpc.Default()
	c.Server.Default()
	c.Database.Default()
	c.Token.Default()
	c.Email.Default()
}
