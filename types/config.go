package types

import (
	"go.heurd.com/heron-go/heron/types/config"
	"reflect"
)

type Config struct {
	ServerPort string
	BeanParsers []BeanParser
	Db struct{
		config.Db
	}
	Application struct{
		config.Application
	}
	Rpc struct {
		config.Rpc
	}
}

type BeanParser struct {

}

func (parser BeanParser) Parse (tag reflect.StructTag, bean reflect.Value, definition reflect.Type) {}
