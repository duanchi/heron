package validate

import (
	"fmt"
	"go.heurd.com/heron-go/heron/types"
	"gopkg.in/go-playground/validator.v9"
	"reflect"
	"strconv"
)

type ValidatorBeanParser struct {
	types.BeanParser
}

func (parser ValidatorBeanParser) Parse (tag reflect.StructTag, bean reflect.Value, definition reflect.Type, beanName string) {
	isValidator, _ := strconv.ParseBool(tag.Get("validator"))

	if isValidator {
		for i := 0; i < definition.NumField(); i++ {
			validateTag := definition.Field(i).Tag.Get("validate-tag")
			validateFunc := definition.Field(i).Tag.Get("validate-func")

			if validateFunc != "" && validateTag != "" {
				if method, has := definition.MethodByName(validateFunc); has {

					Validators[validateTag] = func (fl validator.FieldLevel) bool {

						result := bean.Elem().Method(method.Index).Call([]reflect.Value{
							reflect.ValueOf(fl),
						})

						return result[0].Interface().(bool)
					}
					fmt.Println("Registered " + validateTag + " validator")
				}
			}
		}
	}
}