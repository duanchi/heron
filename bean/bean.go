package bean

import (
	_interface "heurd.com/wand-go/wand/interface"
	"heurd.com/wand-go/wand/server/middleware"
	"heurd.com/wand-go/wand/server/route"
	"reflect"
)

// var beanContainer interface{}
var beanMaps = map[string]reflect.Value{}
// var beanTypeMaps = map[reflect.Type]reflect.Value{}
var beanTypeMaps = map[reflect.Type]reflect.Value{}
var coreBeanParser = []interface{}{
	&route.RouteBeanParser{},
	&middleware.MiddlewareBeanParser{},
}

type Container struct {}

func (bean *Container) Get (name string) reflect.Value {

	beanValue := reflect.ValueOf(bean).Elem()
	beanType := reflect.TypeOf(bean).Elem()

	value := reflect.ValueOf(bean).Elem().FieldByName(name)

	if reflect.Value.IsZero(value) {
		for i := 0; i < beanType.NumField(); i++ {
			if name == beanType.Field(i).Tag.Get("name") {
				return beanValue.Field(i)
			}
		}
	}

	return value
}

func Init(beanContainerInstance interface{}) {
	containerValue := reflect.ValueOf(beanContainerInstance).Elem()
	containerType := reflect.TypeOf(beanContainerInstance).Elem()
	initBean(containerValue, containerType)
}

func initBean(beanContainerInstance reflect.Value, beanContainerType reflect.Type) {
	containerType := beanContainerType
	containerValue := beanContainerInstance
	for i := 0; i < containerValue.NumField(); i++ {
		Register(containerValue.Field(i), containerType.Field(i))
	}

	// beanContainer = beanContainerInstance

	for _, bean := range beanMaps {
		Inject(bean, beanMaps)
	}
}

func Get (name string) interface{} {
	return beanMaps[name].Interface()
}

func GetAll() map[string]reflect.Value {
	return beanMaps
}

func Register (beanValue reflect.Value, beanDefinition reflect.StructField) {
	tag := beanDefinition.Tag
	// beanType := beanDefinition.Type
	name := tag.Get("name")
	if name == "" {
		name = beanDefinition.Name
	}
	beanMaps[name] = reflect.New(beanDefinition.Type).Elem()
	beanTypeMaps[beanMaps[name].Addr().Type()] = beanMaps[name].Addr()

	// beanTypeMaps[beanType] = unsafe.Pointer(beanMaps[name].Addr())

	// beanValue.Set(beanMaps[name].Addr())

	extendParse(tag, beanMaps[name].Addr(), beanDefinition.Type)
}

func extendParse (tag reflect.StructTag, bean reflect.Value, definition reflect.Type) {
	for i := 0; i < len(coreBeanParser); i++ {
		reflect.ValueOf(coreBeanParser[i]).Interface().(_interface.BeanParserInterface).Parse(tag, bean, definition)
	}
}