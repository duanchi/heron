package server

import (
	"wand/core/container"
	_interface "wand/core/interface"
	"wand/core/types"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"reflect"
)

func ExecuteHandle(resource string, ctx *gin.Context, engine *gin.Engine) {
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
			runtimeError := reflect.ValueOf(exception).Interface().(types.Error)
			response.Message = runtimeError.Error()
			statusCode = runtimeError.Code()

			defer func() {
				if err := recover(); err != nil {
					commonError := reflect.ValueOf(exception).Interface().(error)
					response.Message = commonError.Error()
				}

				ctx.JSON(statusCode, response)
			}()
		}
	}()

	// handle := internal.Container.Get(resource).Interface().(_interface.ControllerInterface)
	var data interface{}
	var err error
	bean := container.Container.Get(resource)

	if !bean.IsNil() && !bean.IsZero() {
		handle := bean.Interface().(_interface.ControllerInterface)

		switch method {
		case "GET":
			data, err = handle.Fetch(id, resource, &params, ctx)
		case "POST":
			data, err = handle.Create(id, resource, &params, ctx)
		case "PUT":
			data, err = handle.Update(id, resource, &params, ctx)
		case "DELETE":
			data, err = handle.Remove(id, resource, &params, ctx)
		case "HEAD":
			data, err = handle.Fetch(id, resource, &params, ctx)
		case "OPTIONS":
			data, err = handle.Fetch(id, resource, &params, ctx)
		}
	} else {
		panic(types.RuntimeError{
			Message:   "Resource Not Found",
			ErrorCode: http.StatusNotFound,
		})
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
