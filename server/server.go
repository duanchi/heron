package server

import (
	"github.com/gin-gonic/gin"
	"go.heurd.com/heron-go/heron/config"
	"go.heurd.com/heron-go/heron/server/middleware"
	"go.heurd.com/heron-go/heron/server/route"
	"go.heurd.com/heron-go/heron/server/static"
	"go.heurd.com/heron-go/heron/server/validate"
	"log"
)

var HttpServer *gin.Engine

func Init () {
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
}