package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	_interface "heurd.com/wand-go/wand/interface"
	"heurd.com/wand-go/wand/types"
	"net/http"
	"reflect"
	"runtime/debug"
)

func RestfulHandle(resource string, controller reflect.Value, ctx *gin.Context, engine *gin.Engine) {
	params := ctx.Params
	id := ctx.Param("id")
	method := ctx.Request.Method
	requestId := uuid.NewV4().String()
	response := types.Response{
		RequestId: requestId,
		Status: false,
		Data: nil,
		Message: "Ok",
	}

	defer func() {
		statusCode := http.StatusInternalServerError

		if exception := recover(); exception != nil {
			defer func() {
				if err := recover(); err != nil {
					response.Message = fmt.Sprint(exception)
				}
				debug.PrintStack()
				ctx.JSON(statusCode, response)
			}()
			runtimeError := reflect.ValueOf(exception).Interface().(types.Error)
			fmt.Println(runtimeError.Error())
			response.Message = runtimeError.Error()
			statusCode = runtimeError.Code()
		}
	}()

	var data interface{}
	var err error

	executor := controller.Elem().Interface().(_interface.RestControllerInterface)

	switch method {
	case "GET":
		data, err = executor.Fetch(id, resource, &params, ctx)
	case "POST":
		data, err = executor.Create(id, resource, &params, ctx)
	case "PUT":
		data, err = executor.Update(id, resource, &params, ctx)
	case "DELETE":
		data, err = executor.Remove(id, resource, &params, ctx)
	case "HEAD":
		data, err = executor.Fetch(id, resource, &params, ctx)
	case "OPTIONS":
		data, err = executor.Fetch(id, resource, &params, ctx)
	}

	if err == nil {
		response.Status = true
		response.Data = data
		ctx.JSON(http.StatusOK, response)
	} else {
		panic(err)
	}

	return
}
