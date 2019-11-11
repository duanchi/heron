package types

import "reflect"

type Config struct {
	ServerPort string
	BeanParsers []BeanParser
	Db struct{
		Enabled bool `value:"false"`
		Dsn string
	}
	Rpc struct {
		Server struct{
			Enabled bool `value:"false"`
			Prefix string `value:""`
		}
	}
}

type BeanParser struct {

}

func (parser BeanParser) Parse (tag reflect.StructTag, bean reflect.Value, definition reflect.Type) {}
