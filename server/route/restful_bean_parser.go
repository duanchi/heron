package route

import (
	"go.heurd.com/heron-go/heron/types"
	"reflect"
)

type RestfulBeanParser struct {
	types.BeanParser
}

func (parser RestfulBeanParser) Parse (tag reflect.StructTag, bean reflect.Value, definition reflect.Type, beanName string) {

	resource := tag.Get("restful")

	if resource == "" {
		resource = tag.Get("rest")
	}

	if resource != "" {
		RestfulRoutes[resource] = bean
	}
}