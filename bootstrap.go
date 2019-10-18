package wand

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"heurd.com/wand-go/wand/config"
)
var HttpServer *gin.Engine
var Config config.CoreConfig

func Bootstrap(configuration interface{}, routes map[string]interface{}) {
	config.Init(configuration)
	Config = configuration.(config.CoreConfig)

	fmt.Print(Config.Env)

	/*container.Init()
	server.Init(routes)

	HttpServer = server.HttpServer*/
}