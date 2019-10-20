package wand

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"heurd.com/wand-go/wand/bean"
	"heurd.com/wand-go/wand/config"
)
var HttpServer *gin.Engine

func Bootstrap(configuration interface{}, beanConfiguration interface{}) {
	config.Init(configuration)
	bean.Init(beanConfiguration)

	fmt.Print(config.Get("Application.ServerPort").(string))

	/*bean.Init()
	server.Init(routes)

	HttpServer = server.HttpServer*/
}