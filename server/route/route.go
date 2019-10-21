package route

import (
	"github.com/gin-gonic/gin"
	_interface "heurd.com/wand-go/wand/interface"
	"heurd.com/wand-go/wand/server/handler"
	"reflect"
)

var Routes = map[string]_interface.RestControllerInterface{}

func Init(httpSever *gin.Engine, routes *map[string]_interface.RestControllerInterface) {

	for resource, controller := range reflect.ValueOf(routes).Elem().Interface().(map[string]_interface.RestControllerInterface) {
		httpSever.Any("/" + resource + "/", func(ctx *gin.Context) {
			handler.RestfulHandle(resource, controller, ctx, httpSever)
		})
		httpSever.Any("/" + resource + "/:id", func(ctx *gin.Context) {
			handler.RestfulHandle(resource, controller, ctx, httpSever)
		})
	}
}