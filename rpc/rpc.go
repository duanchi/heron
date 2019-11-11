package rpc

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"heurd.com/wand-go/wand/config"
	"heurd.com/wand-go/wand/server/middleware"
	"heurd.com/wand-go/wand/types"
	"reflect"
	"strings"
)

type RpcBeanMap map[string]struct{
	Package string
	Instance reflect.Value
}

var RpcBeans = RpcBeanMap{}

func (this RpcBeanMap) Init (httpServer *gin.Engine) {
	prefix := config.Get("Rpc.Server.Prefix").(string)

	httpServer.POST(prefix + "/*rpc_path", middleware.HandleAfterRoute, func(ctx *gin.Context) {

		pathStack := strings.SplitN(ctx.Request.URL.Path[len(prefix) + 1:], "::", 2)

		beanName := pathStack[0]
		methodName := pathStack[1]

		if bean, ok := RpcBeans[beanName]; ok {

			method := bean.Instance.MethodByName(methodName)

			if method.IsValid() {

				methodType := method.Type()
				parameters := []interface{}{}
				arguments := []reflect.Value{}
				response := []interface{}{}

				for i := 0; i < methodType.NumIn(); i++ {
					parameters = append(parameters, reflect.New(methodType.In(i)).Elem().Interface())
				}
				ctx.BindJSON(&parameters)

				if methodType.NumIn() != len(parameters) {
					panic(types.RuntimeError{
						Message:   `Malformed arguments in Method "#{methodName}" in Class "#{beanName}"`,
						ErrorCode: 400,
					})
				}

				for i := 0; i < methodType.NumIn(); i++ {
					arguments = append(arguments, reflect.ValueOf(parameters[i]))
				}

				returns := method.Call(arguments)

				for i := 0; i < methodType.NumOut(); i++ {
					response = append(response, returns[i].Interface())
				}

				fmt.Println(response)

				ctx.JSON(200, response)

			} else {
				panic(types.RuntimeError{
					Message:   `No implement of Method "#{methodName}" in Class "#{beanName}"`,
					ErrorCode: 404,
				})
			}
		} else {
			panic(types.RuntimeError{
				Message:   `No implement of Class "#{beanName}"`,
				ErrorCode: 404,
			})
		}
	})
}