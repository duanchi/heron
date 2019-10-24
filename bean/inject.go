package bean

import (
	"fmt"
	"heurd.com/wand-go/wand/config"
	"reflect"
	"regexp"
	"strings"
)

func Inject (rawBean reflect.Value, beanMap map[string]reflect.Value) {

	fmt.Println("Goes !!!!!!!!!")

	beanType := rawBean.Type()

	for i := 0; i < beanType.NumField(); i++ {
		// fieldType := beanType.Field(i).Type
		fieldTag := beanType.Field(i).Tag
		value := fieldTag.Get("value")
		if value != "" {
			regex, _ := regexp.Compile("(" + regexp.QuoteMeta("${") + ".+" + regexp.QuoteMeta("}") + ")")

			regex.ReplaceAllFunc([]byte(value), func(match []byte) []byte {

				env := strings.Split(string(match[2:len(match) - 1]), ",")
				var data string
				data = config.Get(env[0]).(string)

				if len(env) > 1 && (data == "") {
					data = env[1]
				}

				return []byte(data)
			})

			fmt.Println(value)
		}
	}
}


func parseTagNamedValue() {

}