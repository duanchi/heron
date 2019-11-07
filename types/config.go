package types

import "reflect"

type Config struct {
	ServerPort string
	BeanParsers []BeanParser
	Db struct{
		Enabled bool `value:"false"`
		Dsn string
	}
	RpcServer struct{
		Enabled bool `value:"false"`
	}
}

type BeanParser struct {

}

func (parser BeanParser) Parse (tag reflect.StructTag, bean reflect.Value, definition reflect.Type) {}
