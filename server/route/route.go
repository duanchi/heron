package route

import (
	"github.com/gin-gonic/gin"
	"heurd.com/wand-go/wand/server/handler"
	"reflect"
)

var Routes = map[string]reflect.Value{}

func Init(httpSever *gin.Engine) {

	for key, _ := range Routes {

		resource := key

		httpSever.Any("/" + resource + "/", func(ctx *gin.Context) {
			handler.RestfulHandle(resource, Routes[resource], ctx, httpSever)
		})
		httpSever.Any("/" + resource + "/:id", func(ctx *gin.Context) {
			handler.RestfulHandle(resource, Routes[resource], ctx, httpSever)
		})
	}
}