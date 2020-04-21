package cache

import (
	"go.heurd.com/heron-go/heron/types"
	"reflect"
)

type MiddlewareBeanParser struct {
	types.BeanParser
}

func (parser MiddlewareBeanParser) Parse (tag reflect.StructTag, bean reflect.Value, definition reflect.Type, beanName string) {

	cacheName := tag.Get("cache")

	if cacheName != "" {
		CacheEngines[cacheName] = bean
	}
}