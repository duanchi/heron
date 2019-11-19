package heron

import (
	"go.heurd.com/heron-go/heron/bean"
	"go.heurd.com/heron-go/heron/config"
	"go.heurd.com/heron-go/heron/db"
	"go.heurd.com/heron-go/heron/server"
)

func Bootstrap(configuration interface{}, beanConfiguration interface{}) {
	config.Init(configuration)
	Config = configuration

	bean.Init(beanConfiguration)

	if config.Get("Db.Enabled").(bool) == true {
		db.Init()
		Db = db.Connection
	}

	server.Init()
	HttpServer = server.HttpServer
}