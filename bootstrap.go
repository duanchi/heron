package heron

import (
	"go.heurd.com/heron-go/heron/bean"
	"go.heurd.com/heron-go/heron/config"
	"go.heurd.com/heron-go/heron/db"
	_interface "go.heurd.com/heron-go/heron/interface"
	"go.heurd.com/heron-go/heron/server"
)

func Bootstrap(configuration interface{}, beanConfiguration interface{}, beanParsers []_interface.BeanParserInterface) {
	config.Init(configuration)
	Config = configuration

	bean.Init(beanConfiguration, beanParsers)

	if config.Get("Db.Enabled").(bool) == true {
		db.Init()
		Db = db.Connection
	}

	server.Init()
	HttpServer = server.HttpServer
}