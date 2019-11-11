package route

import (
	"github.com/gin-gonic/gin"
	"heurd.com/wand-go/wand/rpc"
)

func Init(httpServer *gin.Engine) {
	BaseRoutes.Init(httpServer)
	RestfulRoutes.Init(httpServer)
	rpc.RpcBeans.Init(httpServer)
}