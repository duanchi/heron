package handler

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	_interface "heurd.com/wand-go/wand/interface"
	"heurd.com/wand-go/wand/types"
	"net/http"
	"reflect"
)

func RestfulHandle(resource string, controller _interface.RestControllerInterface, ctx *gin.Context, engine *gin.Engine) {
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

	switch method {
	case "GET":
		data, err = controller.Fetch(id, resource, &params, ctx)
	case "POST":
		data, err = controller.Create(id, resource, &params, ctx)
	case "PUT":
		data, err = controller.Update(id, resource, &params, ctx)
	case "DELETE":
		data, err = controller.Remove(id, resource, &params, ctx)
	case "HEAD":
		data, err = controller.Fetch(id, resource, &params, ctx)
	case "OPTIONS":
		data, err = controller.Fetch(id, resource, &params, ctx)
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
