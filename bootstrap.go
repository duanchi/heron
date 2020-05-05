package heron

import (
	"fmt"
	"go.heurd.com/heron-go/heron/bean"
	"go.heurd.com/heron-go/heron/cache"
	"go.heurd.com/heron-go/heron/config"
	"go.heurd.com/heron-go/heron/db"
	"go.heurd.com/heron-go/heron/feign"
	"go.heurd.com/heron-go/heron/log"
	"go.heurd.com/heron-go/heron/server"
	"go.heurd.com/heron-go/heron/task"
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

	log.Init(config.Get("Log").(config2.Log))
	Log = &log.Log
	if !checkConfigEnabled("Log.Enabled") {
		Log.Enabled(false)
	}

	if checkConfigEnabled("Task.Enabled") {
		Log.Info("Task Enabled!")
		task.Init()
	}

	if checkConfigEnabled("Db.Enabled") {
		db.Init()
		Db = db.Connection
	}

	if checkConfigEnabled("Cache.Enabled") {
		cache.Init()
	}

	if checkConfigEnabled("Feign.Enabled") {
		feign.Init(config.Get("Feign").(config2.Feign))
	}

	if checkConfigEnabled("Task.Enabled") {
		go task.RunAfterInit()
	}

	go server.Init(errs)
	HttpServer = server.HttpServer

	if checkConfigEnabled("Task.Enabled") {
		go task.RunOnStart()
	}


	go func() {
		c := make(chan os.Signal, 2)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	err := <-errs

	if checkConfigEnabled("Task.Enabled") {
		go task.RunOnExit()
	}


	log.Log.Error("%s", err)
}

func SetConfigFile(configFile string) {
	config.SetConfigFile(configFile)
}

func checkConfigEnabled(configStack string) bool {
	return config.Get(configStack).(bool)
}
