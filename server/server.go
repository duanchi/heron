package server

import (
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
	"log"
	"wand"
	"wand/server/route"
)

var DB *xorm.Engine
var HttpServer *gin.Engine

func Init (routes map[string]interface{}) {
	HttpServer = gin.Default()


	if wand.Config.Env == "production" {
		gin.SetMode("release")
	} else {
		gin.SetMode("debug")
	}


	route.Init(HttpServer, routes)

	serverError := HttpServer.Run(":" + wand.Config.ServerPort)

	if serverError != nil {
		log.Fatal(serverError)
	}
}