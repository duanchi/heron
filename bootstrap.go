package wand

import (
	"fmt"
	"heurd.com/wand-go/wand/bean"
	"heurd.com/wand-go/wand/config"
	"heurd.com/wand-go/wand/db"
	"heurd.com/wand-go/wand/server"
)

func Bootstrap(configuration interface{}, beanConfiguration interface{}) {
	config.Init(configuration)
	bean.Init(beanConfiguration)

	fmt.Print(config.Get("Db.Enabled"))

	if config.Get("Db.Enabled").(bool) == true {
		db.Init()
	}

	server.Init()
	/*bean.Init()
	server.Init(routes)

	HttpServer = server.HttpServer*/
}