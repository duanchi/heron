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
	return GetRaw(key).Interface()
}

func GetRaw(key string) reflect.Value {

	keyStack := strings.Split(key, ".")
	value := reflect.ValueOf(configInstance).Elem()

	for i := 0; i < len(keyStack); i++ {

		//fmt.Printf("key: %s, kind %s",keyStack[i], reflect.TypeOf(value).Kind())

		// 调用栈不是末尾, 并且value是可用值, 并且value是基础类型
		if i < len(keyStack) - 1 && value.IsValid() && value.Kind() != reflect.Ptr && value.Kind() != reflect.Struct {
			return reflect.New(value.Type())
		} else {
			if value.Kind() == reflect.Struct {
				if value.FieldByName(keyStack[i]).IsValid() {
					value = value.FieldByName(keyStack[i])
				} else {
					value = reflect.New(value.Type())
				}
			} else if value.Kind() == reflect.Ptr {
				if value.Elem().FieldByName(keyStack[i]).IsValid() {
					value = value.Elem().FieldByName(keyStack[i])
				} else {
					value = reflect.New(value.Elem().Type())
				}
			} else {
				if value.Elem().FieldByName(keyStack[i]).IsZero() || value.Elem().FieldByName(keyStack[i]).IsNil() {
					value = reflect.New(value.Elem().Type())
				} else {
					value = value.Elem().FieldByName(keyStack[i])
				}
			}
		}

		/*if i < len(keyStack) - 1 && reflect.ValueOf(value).IsValid() && reflect.TypeOf(value).Kind() != reflect.Ptr && reflect.TypeOf(value).Kind() != reflect.Struct {
			return nil
		} else {
			if reflect.TypeOf(value).Kind() == reflect.Struct {
				if reflect.ValueOf(value).FieldByName(keyStack[i]).IsValid() {
					value = reflect.ValueOf(value).FieldByName(keyStack[i]).Interface()
				} else {
					value = new(interface{})
				}
			} else if reflect.TypeOf(value).Kind() == reflect.Ptr {
				if reflect.ValueOf(value).Elem().FieldByName(keyStack[i]).IsValid() {
					value = reflect.ValueOf(value).Elem().FieldByName(keyStack[i]).Interface()
				} else {
					value = struct{}{}
				}

			}else {
				if reflect.ValueOf(value).Elem().FieldByName(keyStack[i]).IsZero() {
					switch reflect.TypeOf(value).Kind() {
					case reflect.Int, reflect.Int64:
						value = 0

					case reflect.String:
						value = ""

					case reflect.Float64:
						value = 0

					case reflect.Bool:
						value = false
					}
				} else if reflect.ValueOf(value).Elem().FieldByName(keyStack[i]).IsNil() {
					value = nil
				} else {
					value = reflect.ValueOf(value).Elem().FieldByName(keyStack[i]).Interface()
				}
			}
		}*/
	}

	return value
}

func parseConfig (config interface{}) {

	configType := reflect.TypeOf(config).Elem()
	configValue := reflect.ValueOf(config).Elem()

	for i := 0; i < configValue.NumField(); i++ {
		if configValue.Field(i).CanInterface() && configValue.Field(i).CanSet() {
			if configValue.Field(i).Kind() == reflect.Struct {
				parseConfig(configValue.Field(i).Addr().Interface())
			} else {
				v := ""
				if configType.Field(i).Tag.Get("env") != "" {
					value := strings.Split(configType.Field(i).Tag.Get("env"), ",")
					if value[0] != "" {
						if len(value) > 1 {
							v = util.Getenv(value[0], value[0])
						} else {
							v = util.Getenv(value[0], "")
						}

					} else {
						v = ""
					}

				}
				if v == "" && configType.Field(i).Tag.Get("value") != "" {
					v = configType.Field(i).Tag.Get("value")
				}

				class := configType.Field(i).Type.Kind()
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