package route

import (
	"wand/config/routes"
	"wand/core/container"
	"wand/core/server"
	"github.com/gin-gonic/gin"
)

func Init (httpServer *gin.Engine, route map[string]interface{}) {
	for name, class := range route {
		container.Container.Register(name, class)
	}

	server.RestfulHandle(httpServer, routes.Routes)
}