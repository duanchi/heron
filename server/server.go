package server

import (
	"github.com/gin-gonic/gin"
	"heurd.com/wand-go/wand/config"
	"heurd.com/wand-go/wand/server/middleware"
	"heurd.com/wand-go/wand/server/route"
	"log"
)

// var DB *xorm.Engine
var HttpServer *gin.Engine

func Init () {
	HttpServer = gin.Default()


	if config.Get("Env").(string) == "production" {
		gin.SetMode("release")
	} else {
		gin.SetMode("debug")
	}

	middleware.Init(HttpServer, middleware.BeforeRoute)

	route.Init(HttpServer)

	serverError := HttpServer.Run(":" + config.Get("Application.ServerPort").(string))

	if serverError != nil {
		log.Fatal(serverError)
	}
}