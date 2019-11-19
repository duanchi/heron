package abstract

import (
	"github.com/gin-gonic/gin"
	_interface "go.heurd.com/heron-go/heron/interface"
)

type Middleware struct {
	_interface.MiddlewareInterface
}

func (this *Middleware) BeforeRoute (ctx *gin.Context) {
	ctx.Next()
}

func (this *Middleware) AfterRoute (ctx *gin.Context) {
	ctx.Next()
}

func (this *Middleware) BeforeResponse (ctx *gin.Context) {
	ctx.Next()
}

func (this *Middleware) AfterResponse (ctx *gin.Context) {
	ctx.Next()
}

func (this *Middleware) AfterPanic (ctx *gin.Context) {
	ctx.Next()
}