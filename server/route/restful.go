package route

import (
	"github.com/gin-gonic/gin"
	"heurd.com/wand-go/wand/server/handler"
	"heurd.com/wand-go/wand/server/middleware"
	"reflect"
)

type RestfulRoutesMap map[string]reflect.Value

var RestfulRoutes = RestfulRoutesMap{}

func (this RestfulRoutesMap) Init (httpServer *gin.Engine) {
	for key, _ := range this {

		resource := key

		httpServer.Any("/" + resource + "/", func (ctx *gin.Context) {
			ctx.Set("resource", resource)
			ctx.Next()
		}, middleware.HandleAfterRoute, func(ctx *gin.Context) {
			handler.RestfulHandle(resource, RestfulRoutes[resource], ctx, httpServer)
		})
		httpServer.Any("/" + resource + "/:id", func (ctx *gin.Context) {
			ctx.Set("resource", resource)
			ctx.Next()
		}, middleware.HandleAfterRoute, func(ctx *gin.Context) {
			handler.RestfulHandle(resource, RestfulRoutes[resource], ctx, httpServer)
		})
	}
}