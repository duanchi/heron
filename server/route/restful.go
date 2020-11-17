package route

import (
	"github.com/gin-gonic/gin"
	"github.com/duanchi/heron/server/handler"
	"github.com/duanchi/heron/server/middleware"
	"reflect"
)

type RestfulRoutesMap map[string]reflect.Value

var RestfulRoutes = RestfulRoutesMap{}

func (this RestfulRoutesMap) Init (httpServer *gin.Engine) {
	for key, _ := range this {

		resource := key

		handlers := []gin.HandlerFunc{
			func (ctx *gin.Context) {
				ctx.Set("resource", resource)
				ctx.Next()
			},
		}

		handlers = middleware.GetHandlersAfterRouterAppend(handlers)

		handlers = append(handlers, func(ctx *gin.Context) {
			handler.RestfulHandle(resource, RestfulRoutes[resource], ctx, httpServer)
		})

		httpServer.Any("/" + resource, handlers...)
		httpServer.Any("/" + resource + "/", handlers...)
		httpServer.Any("/" + resource + "/:id", handlers...)
	}
}