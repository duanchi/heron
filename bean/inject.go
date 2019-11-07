package bean

import (
	"heurd.com/wand-go/wand/config"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

func Inject (rawBean reflect.Value, beanMap map[string]reflect.Value) {

	beanType := rawBean.Type()

	for i := 0; i < beanType.NumField(); i++ {
		if rawBean.Field(i).CanSet() {
			fieldTag := beanType.Field(i).Tag

			parseTagNamedValue(fieldTag.Get("value"), rawBean.Field(i))
			parseTagNamedAutowired(fieldTag.Get("autowired"), rawBean.Field(i))
		}

	}
}

func parseTagNamedValue(value string, field reflect.Value) {
	if value != "" {
		regex, _ := regexp.Compile("^" + regexp.QuoteMeta("${") + "(.+)" + regexp.QuoteMeta("}") + "$")

		value = string(regex.ReplaceAllFunc([]byte(value), func(match []byte) []byte {
			return match[2:len(match) - 1]
		})[:])

		configField := strings.Split(value, ",")
		configValue := config.GetRaw(configField[0])

		if configValue.IsZero() && len(configField) > 1 {

			class := field.Kind()

			switch class {
			case reflect.String:
				configValue.SetString(configField[1])

			case reflect.Int, reflect.Int64:
				value, err := strconv.ParseInt(configField[1], 10, 64)
				if err != nil {
					configValue.SetInt(0)
				} else {
					configValue.SetInt(value)
				}

			case reflect.Bool:
				value, err := strconv.ParseBool(configField[1])
				if err != nil {
					configValue.SetBool(false)
				} else {
					configValue.SetBool(value)
				}

			case reflect.Float64:
				value, err := strconv.ParseFloat(configField[1], 10)
				if err != nil {
					configValue.SetFloat(0)
				} else {
					configValue.SetFloat(value)
				}
			}
		} else {
			field.Set(configValue)
		}


	}
}

func parseTagNamedAutowired(value string, field reflect.Value) {
	setAutowired, _ := strconv.ParseBool(value)
	beanType := field.Type()
	if setAutowired && beanType.Kind() == reflect.Ptr {
		beanPointer, ok := beanTypeMaps[beanType]
		if ok {
			field.Set(beanPointer)
		}
	}
}