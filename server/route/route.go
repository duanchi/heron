package route

import (
	"github.com/gin-gonic/gin"
	"go.heurd.com/heron-go/heron/config"
	"go.heurd.com/heron-go/heron/rpc"
)

func Init(httpServer *gin.Engine) {
	BaseRoutes.Init(httpServer)
	RestfulRoutes.Init(httpServer)
	if config.Get("Rpc.Server.Enabled").(string) == "true" {
		rpc.RpcBeans.Init(httpServer)
	}
}