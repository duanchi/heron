package handler

import (
	"github.com/gin-gonic/gin"
	_interface "github.com/duanchi/heron/interface"
	"reflect"
)

func RouteHandle(path string, handle reflect.Value, ctx *gin.Context, engine *gin.Engine) {
	params := ctx.Params
	method := ctx.Request.Method

	handle.Interface().(_interface.RouterInterface).Handle(ctx.Request.URL.Path, method, params, ctx)

	return
}
