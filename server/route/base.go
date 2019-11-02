package route

import (
	"github.com/gin-gonic/gin"
	"heurd.com/wand-go/wand/server/handler"
	"heurd.com/wand-go/wand/server/middleware"
	"reflect"
	"strings"
)

type BaseRoutesMap map[string]reflect.Value

var BaseRoutes = BaseRoutesMap{}

func (THIS BaseRoutesMap) Init (httpServer *gin.Engine) {
	for key, _ := range THIS {

		name := key

		stack := strings.SplitN(name, "@", 2)
		route := "/"
		methods := []string{"ALL"}

		if stack[0] != "" {
			route = stack[0]
		}
		if len(stack) > 1 && stack[1] != "" {
			methods = strings.Split(strings.ToUpper(stack[1]), ",")
		}

		for _, method := range methods {

			if method == "ALL" {
				httpServer.Any(route, middleware.HandleAfterRoute, func(ctx *gin.Context) {
					handler.RouteHandle(route, BaseRoutes[name], ctx, httpServer)
				})
			} else {
				httpServer.Handle(method, route, middleware.HandleAfterRoute, func(ctx *gin.Context) {
					handler.RouteHandle(route, BaseRoutes[name], ctx, httpServer)
				})
			}
		}
	}
}