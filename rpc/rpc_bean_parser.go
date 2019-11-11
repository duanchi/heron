package rpc

import (
	"heurd.com/wand-go/wand/types"
	"reflect"
)

type RpcBeanParser struct {
	types.BeanParser
}

func (parser RpcBeanParser) Parse (tag reflect.StructTag, bean reflect.Value, definition reflect.Type, beanName string) {

	rpc := tag.Get("rpc")

	if rpc != "" {
		RpcBeans[bean.Elem().Type().PkgPath() + "." + bean.Elem().Type().Name()] = bean
	}
}