package rpc

import (
	"bytes"
	"encoding/json"
	_interface "heurd.com/wand-go/wand/interface"
	"heurd.com/wand-go/wand/types"
	"io/ioutil"
	"net/http"
	"reflect"
	"runtime"
	"strings"
)

func Call (in IN, out *OUT, caller _interface.RpcInterface) (err error) {



	caller.GetPackageName()



	pc, _, _, _ := runtime.Caller(1)

	methodStack := strings.Split(runtime.FuncForPC(pc).Name(), ".")

	*out, err = rpcRequest(
		caller.GetApplicationName(),
		caller.GetPackageName(),
		reflect.ValueOf(caller).Elem().Type().String(),
		methodStack[len(methodStack) - 1],
		&in,
	)

	return
}

func rpcRequest(serviceName string, packageName string, className string, method string, in *IN) (out OUT, err error) {
	client := &http.Client{}

	requestBody, _ := json.Marshal(in)

	request, err := http.NewRequest(http.MethodPost, "http://" + serviceName + "/" + packageName + "." + className + "::" + method, bytes.NewReader(requestBody))
	if err != nil {
		// handle error
		err = types.RuntimeError{
			Message:   "RPC request error",
			ErrorCode: 500,
		}
		return
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Cookie", "name=anny")

	response, err := client.Do(request)

	if err != nil {
		err = types.RuntimeError{
			Message:   "RPC response error",
			ErrorCode: 500,
		}
		return
	}

	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)

	err = json.Unmarshal(responseBody, &out)

	return
}