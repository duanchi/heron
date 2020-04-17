package heron

import (
	"fmt"
	"go.heurd.com/heron-go/heron/bean"
	"go.heurd.com/heron-go/heron/config"
	"go.heurd.com/heron-go/heron/db"
	"go.heurd.com/heron-go/heron/feign"
	"go.heurd.com/heron-go/heron/log"
	"go.heurd.com/heron-go/heron/server"
	config2 "go.heurd.com/heron-go/heron/types/config"
	"os"
	"os/signal"
	"syscall"
)

func Bootstrap(configuration interface{}) {
	config.Init(configuration)
	Config = configuration

	errs := make(chan error, 3)

	bean.Init(
		config.Get("Beans"),
		config.Get("BeanParsers"),
	)

	if config.Get("Db.Enabled").(bool) {
		db.Init()
		Db = db.Connection
	}

	if config.Get("Log.Enabled").(bool) {
		log.Init(config.Get("Log").(config2.Log))
		Log = &log.Log
	}


	feign.Init(config.Get("Feign").(config2.Feign))

	go server.Init(errs)
	HttpServer = server.HttpServer

	go func() {
		c := make(chan os.Signal, 2)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	err := <-errs

	log.Log.Error("%s", err)
}

func SetConfigFile(configFile string) {
	config.SetConfigFile(configFile)
}
