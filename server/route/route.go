package route

import (
	"github.com/gin-gonic/gin"
)

func Init(httpServer *gin.Engine) {
	BaseRoutes.Init(httpServer)
	RestfulRoutes.Init(httpServer)
}