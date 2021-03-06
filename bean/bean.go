package bean

import (
	"fmt"
	_interface "go.heurd.com/heron-go/heron/interface"
	"go.heurd.com/heron-go/heron/rpc"
	"go.heurd.com/heron-go/heron/server/middleware"
	"go.heurd.com/heron-go/heron/server/route"
	"go.heurd.com/heron-go/heron/service"
	"reflect"
)

// var beanContainer interface{}
var beanMaps = map[string]reflect.Value{}
var beanNameMaps = map[string]reflect.Value{}
var beanTypeMaps = map[reflect.Type]reflect.Value{}
var coreBeanParsers = []_interface.BeanParserInterface{
	&service.ServiceBeanParser{},
	&route.RouteBeanParser{},
	&route.RestfulBeanParser{},
	&middleware.MiddlewareBeanParser{},

	&rpc.RpcBeanParser{},
}

var customBeanParsers = []_interface.BeanParserInterface{}

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

func Init(beanContainerInstance interface{}, beanParsers interface{}) {

	if reflect.ValueOf(beanParsers).IsValid() {
		customBeanParsers = beanParsers.([]_interface.BeanParserInterface)
	}

	containerValue := reflect.ValueOf(beanContainerInstance)
	containerType := reflect.TypeOf(beanContainerInstance)

	if reflect.TypeOf(beanContainerInstance).Kind() == reflect.Ptr {
		containerValue = reflect.ValueOf(beanContainerInstance).Elem()
		containerType = reflect.TypeOf(beanContainerInstance).Elem()
	}

	initBean(containerValue, containerType)
}

func initBean(beanContainerInstance reflect.Value, beanContainerType reflect.Type) {
	containerType := beanContainerType
	containerValue := beanContainerInstance
	for i := 0; i < containerValue.NumField(); i++ {
		Register(containerValue.Field(i), containerType.Field(i))
	}

	for _, bean := range beanMaps {
		Inject(bean, beanMaps)
	}

	for _, bean := range beanMaps {
		parseInit(bean)
	}
}

func Get (name string) interface{} {
	return beanNameMaps[name].Interface()
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
	beanNameMaps[name] = beanMaps[name].Addr()
	beanTypeMaps[beanMaps[name].Addr().Type()] = beanMaps[name].Addr()

	extendParse(tag, beanMaps[name].Addr(), beanDefinition.Type, name)

	fmt.Println("[Wand-Go] Init Bean: " + name)
}

func extendParse (tag reflect.StructTag, bean reflect.Value, definition reflect.Type, beanName string) {
	for i := 0; i < len(coreBeanParsers); i++ {
		reflect.ValueOf(coreBeanParsers[i]).Interface().(_interface.BeanParserInterface).Parse(tag, bean, definition, beanName)
	}

	for i := 0; i < len(customBeanParsers); i++ {
		reflect.ValueOf(customBeanParsers[i]).Interface().(_interface.BeanParserInterface).Parse(tag, bean, definition, beanName)
	}
}

func parseInit(rawBean reflect.Value) {
	defer func() {
		if err := recover(); err != nil {
		}
	}()
	rawBean.Addr().Interface().(_interface.BeanInterface).Init()
}