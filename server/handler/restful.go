package handler

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	_interface "go.heurd.com/heron-go/heron/interface"
	"go.heurd.com/heron-go/heron/server/websocket"
	"go.heurd.com/heron-go/heron/types"
	"net/http"
	"reflect"
	"runtime/debug"
	"strings"
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
				/*if err := recover(); err != nil {
					response.Message = fmt.Sprint(exception)
				}
				if ctx.Writer.Status() != http.StatusOK {
					statusCode = ctx.Writer.Status()
				}
				ctx.Writer.Header().Set("Content-Type", "application/json")*/
				if ctx.Writer.Status() != http.StatusOK {
					statusCode = ctx.Writer.Status()
				}
				ctx.JSON(statusCode, response)
				debug.PrintStack()
			}()


			_, implemented := exception.(types.Error)

			if implemented {
				runtimeError := reflect.ValueOf(exception).Interface().(types.Error)
				response.Message = runtimeError.Error()
				statusCode = runtimeError.Code()
				response.Data = runtimeError.Data()
			} else {
				commonError := reflect.ValueOf(exception).Interface().(error)
				response.Message = commonError.Error()
			}
		}
	}()

	var data interface{}
	var err error

	executor := controller.Interface().(_interface.RestControllerInterface)

	// Upgrade Protocol to Websocket
	if method == "GET" {
		upgradeRequest := ctx.Request.Header.Get("Connection")
		upgradeProtocol := ctx.Request.Header.Get("Upgrade")

		if upgradeRequest == "Upgrade" && strings.ToLower(upgradeProtocol) == "websocket" {
			websocket.Handle(id, resource, &params, ctx, executor.Communicate)
			return
		}
	}

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
