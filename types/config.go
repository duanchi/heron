package types

import (
	"go.heurd.com/heron-go/heron/types/config"
	"reflect"
)

type Config struct {
	ServerPort string
	BeanParsers []interface{}
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

func (this *Config) GetName() (name string) {
	return "Config"
}
func (this *Config) SetName(name string) {
	return
}

type BeanParser struct {

}

func (parser BeanParser) Parse (tag reflect.StructTag, bean reflect.Value, definition reflect.Type) {}
