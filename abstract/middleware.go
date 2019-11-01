package abstract

import (
	"github.com/gin-gonic/gin"
)

type Middleware struct {}

func (THIS *Middleware) BeforeRoute (ctx *gin.Context) {
	ctx.Next()
}

func (THIS *Middleware) AfterRoute (ctx *gin.Context) {
	ctx.Next()
}

func (THIS *Middleware) BeforeResponse (ctx *gin.Context) {
	ctx.Next()
}

func (THIS *Middleware) AfterResponse (ctx *gin.Context) {
	ctx.Next()
}

func (THIS *Middleware) AfterPanic (ctx *gin.Context) {
	ctx.Next()
}