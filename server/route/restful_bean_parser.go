package route

import (
	"heurd.com/wand-go/wand/types"
	"reflect"
)

type RestfulBeanParser struct {
	types.BeanParser
}

func (parser RestfulBeanParser) Parse (tag reflect.StructTag, bean reflect.Value, definition reflect.Type) {

	resource := tag.Get("restful")

	if resource == "" {
		resource = tag.Get("rest")
	}

	if resource != "" {
		RestfulRoutes[resource] = bean
	}
}