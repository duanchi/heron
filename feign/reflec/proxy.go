package reflec

import (
	"errors"
	"fmt"
	"github.com/duanchi/heron/common"
	"reflect"
)

type Proxy struct {
target  interface{}
methods map[string]*Method
handle  InvocationHandler
}


func New(target interface{}, h InvocationHandler) *Proxy {
typ := reflect.TypeOf(target)
value := reflect.ValueOf(target)
methods := make(map[string]*Method, 0)
for i := 0; i < value.NumMethod(); i++ {
method := value.Method(i)
methods[typ.Method(i).Name] = &Method{value:method}
}
return &Proxy{target:target, methods:methods, handle:h}
}

func (p * Proxy) Get(serviceName string, path string, obj interface{}) (error) {
	return p.GetWithParams(serviceName, path, nil, obj)
}

func (p * Proxy) GetWithParams(serviceName string, path string, params map[string]string, resp interface{}) (error) {
	return p.InvokeMethod(common.Get.String(), serviceName, path, params, resp)
}

func (p * Proxy) Post(serviceName string, path string, body interface{}, resp interface{}) (error) {
	return p.PostWithParams(serviceName, path, nil, body, resp)
}

func (p * Proxy) PostWithParams(serviceName string, path string, params map[string]string, body interface{}, resp interface{}) (error) {
	return p.InvokeMethod(common.Post.String(), serviceName, path, params, body, resp)
}

func (p * Proxy) Put(serviceName string, path string, body interface{}, resp interface{}) (error) {
	return p.PutWithParams(serviceName, path, nil, body, resp)
}

func (p * Proxy) PutWithParams(serviceName string, path string, params map[string]string, body interface{}, resp interface{}) (error) {
	return p.InvokeMethod(common.Put.String(), serviceName, path, params, body, resp)
}

func (p * Proxy) Delete(serviceName string, path string, resp interface{}) (error) {
	return p.DeleteWithParams(serviceName, path, nil, resp)
}

func (p * Proxy) DeleteWithParams(serviceName string, path string, params map[string]string, resp interface{}) (error) {
	return p.InvokeMethod(common.Delete.String(), serviceName, path, params, resp)
}

func (p *Proxy)InvokeMethod(name string, args ...interface{}) (error) {
return p.handle.Invoke(p, p.methods[name], args)
}


type Method struct {
value reflect.Value
}

func (m *Method)Invoke(args ...interface{}) (error){
defer func() {
if p := recover(); p != nil {
err := errors.New(fmt.Sprintf("%s", p))
fmt.Println(err)
}
}()
params := make([]reflect.Value, 0)
if args != nil {
for i := 0; i < len(args); i++ {
params = append(params, reflect.ValueOf(args[i]))
}
}
r := m.value.Call(params)
if r != nil && len(r) > 0{
	err,ok := r[0].Interface().(error)
	if !ok {
		err = nil
	}
	return err
}
return nil
}