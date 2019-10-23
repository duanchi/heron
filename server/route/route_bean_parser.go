package route

import (
	"heurd.com/wand-go/wand/types"
	"reflect"
)

type RouteBeanParser struct {
	types.BeanParser
}

func (parser RouteBeanParser) Parse (tag reflect.StructTag, bean reflect.Value, definition reflect.Type) {

	resource := tag.Get("route")

	if resource != "" {
		Routes[resource] = bean
	}
}