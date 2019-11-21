package service

import (
	_interface "go.heurd.com/heron-go/heron/interface"
	"go.heurd.com/heron-go/heron/types"
	"reflect"
)

type ServiceBeanParser struct {
	types.BeanParser
}

func (parser ServiceBeanParser) Parse (tag reflect.StructTag, bean reflect.Value, definition reflect.Type, beanName string) {
	if definition.Implements(reflect.TypeOf((*_interface.ServiceInterface)(nil)).Elem()) {
		bean.Elem().Interface().(_interface.ServiceInterface).SetServiceName(beanName)
	}
}