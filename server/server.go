package server

import (
	"github.com/gin-gonic/gin"
	"github.com/duanchi/heron/config"
	"github.com/duanchi/heron/server/middleware"
	"github.com/duanchi/heron/server/route"
	"github.com/duanchi/heron/server/static"
	"github.com/duanchi/heron/server/validate"
	"log"
)

var HttpServer *gin.Engine

func Init (err chan error) {
	HttpServer = gin.Default()


	if config.Get("Env").(string) == "production" {
		gin.SetMode("release")
	} else {
		gin.SetMode("debug")
	}

	validate.Init()

	middleware.Init(HttpServer, middleware.BeforeRoute)

	static.Init(HttpServer)

	route.Init(HttpServer)

	serverError := HttpServer.Run(":" + config.Get("Application.ServerPort").(string))

	if serverError != nil {
		log.Fatal(serverError)
	}

	err <- serverError

	return
}