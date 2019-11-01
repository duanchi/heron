package middleware

import (
	"heurd.com/wand-go/wand/types"
	"reflect"
	"strconv"
)

type MiddlewareBeanParser struct {
	types.BeanParser
}

func (parser MiddlewareBeanParser) Parse (tag reflect.StructTag, bean reflect.Value, definition reflect.Type) {

	isMiddleware, _ := strconv.ParseBool(tag.Get("middleware"))

	if isMiddleware {
		Middlewares = append(Middlewares, bean)
	}
}