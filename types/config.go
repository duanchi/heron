package types

import (
	"go.heurd.com/heron-go/heron/types/config"
	"reflect"
)

type Config struct {
	BeanParsers []interface{}
	Db			config.Db `yaml:"db"`
	Application config.Application `yaml:"application"`
	Rpc 		config.Rpc `yaml:"rpc"`
	Feign		config.Feign `yaml:"feign"`
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
