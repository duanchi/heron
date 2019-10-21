package wand

import (
	"heurd.com/wand-go/wand/bean"
	"heurd.com/wand-go/wand/config"
	"heurd.com/wand-go/wand/db"
	"heurd.com/wand-go/wand/server"
)

func Bootstrap(configuration interface{}, beanConfiguration interface{}) {
	config.Init(configuration)
	Config = configuration

	bean.Init(beanConfiguration)
	Bean = beanConfiguration

	if config.Get("Db.Enabled").(bool) == true {
		db.Init()
		Db = db.Connection
	}

	server.Init()
	HttpServer = server.HttpServer
}