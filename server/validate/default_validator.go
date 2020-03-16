package validate

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator"
	"reflect"
	"sync"
)

type defaultValidator struct {
	once     sync.Once
	validate *validator.Validate
}

var _ binding.StructValidator = &defaultValidator{}

func (v *defaultValidator) ValidateStruct(obj interface{}) error {

	if kindOfData(obj) == reflect.Struct {

		v.lazyInit()

		if err := v.validate.Struct(obj); err != nil {

			objType := reflect.TypeOf(reflect.ValueOf(obj).Elem().Interface())

			fmt.Println(objType)

			for n := 0; n < objType.NumField(); n++ {
				fmt.Println(objType.Field(n).Tag.Get("validate"), objType.Field(n).Tag.Get("comment"))
			}

			fmt.Printf("%v", err)

			for _, err := range err.(validator.ValidationErrors) {
				fmt.Println(err.ActualTag())
			}

			return err
		}
	}

	return nil
}

func (v *defaultValidator) Engine() interface{} {
	v.lazyInit()
	return v.validate
}

func (v *defaultValidator) lazyInit() {
	v.once.Do(func() {
		v.validate = validator.New()
		v.validate.SetTagName("binding")
		v.validate.SetTagName("validate")

		// add any custom validations etc. here
	})
}

func kindOfData(data interface{}) reflect.Kind {

	value := reflect.ValueOf(data)
	valueType := value.Kind()

	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	return valueType
}