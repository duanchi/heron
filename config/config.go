package config

import (
	"heurd.com/wand-go/wand/util"
	"reflect"
	"strings"
)

type ConfigInterface interface {
	Get()
}

type CoreConfig struct {
	ServerPort string
	Env string
}

func (config *CoreConfig) Get (key string) string {
	return reflect.ValueOf(config).Elem().FieldByName(key).String()
}

func Init(config interface{}) {
	parseConfig(config)
}

func parseConfig (config interface{}) {

	configType := reflect.TypeOf(config).Elem()
	configValue := reflect.ValueOf(config).Elem()

	for i := 0; i < configValue.NumField(); i++ {
		if configValue.Field(i).CanInterface() && configValue.Field(i).CanSet() {
			if configValue.Field(i).Kind() == reflect.Struct {
				parseConfig(configValue.Field(i).Addr().Interface())
			} else {
				value := strings.Split(configType.Field(i).Tag.Get("env"), ",")
				if value[0] != "" {
					v := util.Getenv(value[0], value[1])
					configValue.Field(i).SetString(v)
				}
			}
		}
	}
}