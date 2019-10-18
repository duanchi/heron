package container

import (
	"wand/core/util"
	"reflect"
)

var Container = container{}

type container map[string]reflect.Value

func (bean container) Register (name string, instance interface{}) {
	bean[name] = reflect.New(util.GetType(instance))
}

func (bean container) Get (name string) reflect.Value {
	return bean[name]
}

func Init() {
}