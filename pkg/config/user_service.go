package config

import (
	"stuoj-common/pkg/config"
)

type UserServiceConf struct {
	Grpc     config.GrpcConf     `yaml:"grpc" json:"grpc"`
	Database config.DatabaseConf `yaml:"database" json:"database"`
	Email    EmailConf           `yaml:"email" json:"email"`
}

func (c *UserServiceConf) Default() {
	c.Grpc.Default()
	c.Database.Default()
	c.Email.Default()
}
