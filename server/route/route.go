package route

import (
	"github.com/gin-gonic/gin"
	"github.com/duanchi/heron/config"
	"github.com/duanchi/heron/rpc"
)

func Init(httpServer *gin.Engine) {
	BaseRoutes.Init(httpServer)
	RestfulRoutes.Init(httpServer)
	if config.Get("Rpc.Server.Enabled").(bool) == true {
		rpc.RpcBeans.Init(httpServer)
	}
}