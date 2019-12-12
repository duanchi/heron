package types

import (
	"go.heurd.com/heron-go/heron/types/config"
	"reflect"
)

type Config struct {
	BeanParsers []interface{}
	Db struct{
		config.Db `yaml:",inline"`
	} `yaml:"db"`
	Application struct{
		config.Application `yaml:",inline"`
	} `yaml:"application"`
	Rpc struct {
		config.Rpc `yaml:",inline"`
	} `yaml:"rpc"`
	Feign struct{
		config.Feign `yaml:",inline"`
	} `yaml:"feign"`
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
