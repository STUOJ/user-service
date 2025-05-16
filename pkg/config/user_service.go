package config

import "common/pkg/config"

type UserServiceConf struct {
	Grpc     config.GrpcConf     `yaml:"grpc" json:"grpc"`
	Database config.DatabaseConf `yaml:"database" json:"database"`
	Email    EmailConf           `yaml:"email" json:"email"`
	Token    TokenConf           `yaml:"token" json:"token"`
}

func (c *UserServiceConf) Default() {
	c.Grpc.Default()
	c.Database.Default()
	c.Token.Default()
	c.Email.Default()
}
