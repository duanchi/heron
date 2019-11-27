package heron

import (
	"go.heurd.com/heron-go/heron/bean"
	"go.heurd.com/heron-go/heron/config"
	"go.heurd.com/heron-go/heron/db"
	"go.heurd.com/heron-go/heron/feign"
	_interface "go.heurd.com/heron-go/heron/interface"
	"go.heurd.com/heron-go/heron/server"
	"go.heurd.com/heron-go/heron/yconfig"
)

func Bootstrap(configuration interface{}, beanConfiguration interface{}, beanParsers []_interface.BeanParserInterface) {
	config.Init(configuration)
	Config = configuration

	bean.Init(beanConfiguration, beanParsers)

	if config.Get("Db.Enabled").(string) == "true" {
		db.Init()
		Db = db.Connection
	}

	server.Init()
	HttpServer = server.HttpServer
}

func BootstrapWithYaml(configuration interface{}, beanConfiguration interface{}, beanParsers []_interface.BeanParserInterface) {
	yconfig,_ := yconfig.GetYamlConfig()
	config.SetconfigInstance(yconfig)

	bean.Init(beanConfiguration, beanParsers)

	if config.Get("Db.Enabled").(string) == "true" {
		db.Init()
		Db = db.Connection
	}

	feign.Init(yconfig.Feign)

	server.Init()
	HttpServer = server.HttpServer
}