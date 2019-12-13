package heron

import (
	"go.heurd.com/heron-go/heron/bean"
	"go.heurd.com/heron-go/heron/config"
	"go.heurd.com/heron-go/heron/db"
	"go.heurd.com/heron-go/heron/feign"
	"go.heurd.com/heron-go/heron/server"
	config2 "go.heurd.com/heron-go/heron/types/config"
)

func Bootstrap(configuration interface{}) {
	config.Init(configuration)
	Config = configuration

	bean.Init(config.Get("Beans"), config.Get("BeanParsers"))

	if config.Get("Db.Enabled").(bool) == true {
		db.Init()
		Db = db.Connection
	}

	feign.Init(config.Get("Feign").(config2.Feign))

	server.Init()
	HttpServer = server.HttpServer
}