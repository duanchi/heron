package yconfig

import (
	"log"
	"os"
	"reflect"
	"strings"
)


func getRaw(key string) reflect.Value {

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
			} else if configValue.Field(i).Kind() == reflect.Slice {
				for j:=0; j < configValue.Field(i).Len(); j++ {
					parseConfig(configValue.Field(i).Index(j).Addr().Interface())
				}
			} else {
				defaultValue := configType.Field(i).Tag.Get("default")
				yamlValue := configValue.Field(i).String()
				//未配置则取默认值
				if yamlValue == "" && defaultValue != "" {
					configValue.Field(i).SetString(defaultValue)
				}
				envValue := ""
				envKey := ""
				value := ""
				if strings.Index(yamlValue, "$") != -1 {
					start := strings.Index(yamlValue, "{")
					end := strings.LastIndex(yamlValue, "}")
					elContent := yamlValue[start:end]
					index := strings.Index(elContent, ":")
					if index != -1 {
						envKey = elContent[1:index]
						value = elContent[index+1:]
					} else {//不存在配置值
						envKey = elContent[1:]
					}

					envValue = os.Getenv(envKey)
					if envValue != "" {
						log.Println(envKey + ": " + envValue)
					}
					//fmt.Println(envKey + ": " + envValue)
					//环境变量不存在则获取缺省值
					if envValue == "" && value != "" {
						envValue = value
					}
					if envValue != "" {
						configValue.Field(i).SetString(envValue)
					}
				}
			}
		}
	}
}
