package config

import (
	"heurd.com/wand-go/wand/util"
	"reflect"
	"strconv"
	"strings"
)

var configInstance interface{}

func Init(config interface{}) {
	parseConfig(config)
	configInstance = config
}

func Get(key string) interface{} {

	keyStack := strings.Split(key, ".")
	value := configInstance
	isStruct := false
	for i := 0; i < len(keyStack); i++ {

		if reflect.TypeOf(value).Kind() == reflect.Struct {
			isStruct = true
		}

		if i < len(keyStack) - 1 && reflect.TypeOf(value).Kind() != reflect.Ptr && reflect.TypeOf(value).Kind() != reflect.Struct {
			return nil
		} else {
			if isStruct {
				value = reflect.ValueOf(value).FieldByName(keyStack[i]).Interface()
			} else {
				value = reflect.ValueOf(value).Elem().FieldByName(keyStack[i]).Interface()
			}
		}
	}

	return reflect.ValueOf(value).Interface()
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
				class := configType.Field(i).Type.Kind()
				if value[0] != "" {
					v := util.Getenv(value[0], value[1])
					switch class {
					case reflect.String:
						configValue.Field(i).SetString(v)

					case reflect.Int, reflect.Int64:
						value, err := strconv.ParseInt(v, 10, 64)
						if err != nil {
							configValue.Field(i).SetInt(0)
						} else {
							configValue.Field(i).SetInt(value)
						}

					case reflect.Bool:
						value, err := strconv.ParseBool(v)
						if err != nil {
							configValue.Field(i).SetBool(false)
						} else {
							configValue.Field(i).SetBool(value)
						}

					case reflect.Float64:
						value, err := strconv.ParseFloat(v, 10)
						if err != nil {
							configValue.Field(i).SetFloat(0)
						} else {
							configValue.Field(i).SetFloat(value)
						}
					}
				}
			}
		}
	}
}