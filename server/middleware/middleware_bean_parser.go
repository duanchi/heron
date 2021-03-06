package middleware

import (
	"go.heurd.com/heron-go/heron/types"
	"reflect"
	"strconv"
)

type MiddlewareBeanParser struct {
	types.BeanParser
}

func (parser MiddlewareBeanParser) Parse (tag reflect.StructTag, bean reflect.Value, definition reflect.Type, beanName string) {

	isMiddleware, _ := strconv.ParseBool(tag.Get("middleware"))

	if isMiddleware {
		Middlewares = append(Middlewares, bean)
	}
}