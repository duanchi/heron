package handler

import (
	"github.com/gin-gonic/gin"
	_interface "heurd.com/wand-go/wand/interface"
	"reflect"
)

func RouteHandle(path string, handle reflect.Value, ctx *gin.Context, engine *gin.Engine) {
	params := ctx.Params
	method := ctx.Request.Method

	handle.Interface().(_interface.RouterInterface).Handle(path, method, params, ctx)

	return
}
