package route

import (
	"github.com/gin-gonic/gin"
	"heurd.com/wand-go/wand/server"
	"heurd.com/wand-go/wand/types"
	"reflect"
)

type RouteBeanParser struct {
	types.BeanParser
}

func (parser RouteBeanParser) Parse (tag reflect.StructTag, bean reflect.Value, definition reflect.Type) {

	resource := tag.Get("route")

	if resource != "" {
		server.HttpServer.Any("/" + resource + "/", func(ctx *gin.Context) {
			server.ExecuteHandle(resource, ctx, server.HttpServer)
		})
		server.HttpServer.Any("/" + resource + "/:id", func(ctx *gin.Context) {
			server.ExecuteHandle(resource, ctx, server.HttpServer)
		})
	}
}