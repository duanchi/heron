package types

import "reflect"

type Config struct {
	ServerPort string
	BeanParsers []BeanParser
}

type BeanParser struct {

}

func (parser BeanParser) Parse (tag reflect.StructTag, bean reflect.Value, definition reflect.Type) {}
