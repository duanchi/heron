package route

import (
	"github.com/gin-gonic/gin"
	"go.heurd.com/heron-go/heron/server/handler"
	"go.heurd.com/heron-go/heron/server/middleware"
	"reflect"
	"strings"
)

type BaseRoutesMap map[string]reflect.Value

var BaseRoutes = BaseRoutesMap{}

func (this BaseRoutesMap) Init (httpServer *gin.Engine) {
	for key, _ := range this {

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

			handlers := middleware.GetHandlersAfterRouter()

			handlers = append(handlers, func(ctx *gin.Context) {
				handler.RouteHandle(route, BaseRoutes[name], ctx, httpServer)
			})

			if method == "ALL" {
				httpServer.Any(route, handlers...)
			} else {
				httpServer.Handle(method, route, handlers...)
			}
		}
	}
}