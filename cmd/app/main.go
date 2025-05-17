package app

import "user-service/cmd/bootstrap"

func Main() {
	bootstrap.InitConfig()
	bootstrap.InitNacos()
	bootstrap.InitDatabase()
	//bootstrap.InitRpc()
	bootstrap.InitServer()
}
