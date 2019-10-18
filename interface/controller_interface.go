package _interface

import (
	"wand/core/types"
	"github.com/gin-gonic/gin"
)

type ControllerInterface interface {
	Fetch(id string, resource string, parameters *gin.Params, ctx *gin.Context) (result interface{}, err types.Error)

	Create(id string, resource string, parameters *gin.Params, ctx *gin.Context) (result interface{}, err types.Error)

	Update(id string, resource string, parameters *gin.Params, ctx *gin.Context) (result interface{}, err types.Error)

	Remove(id string, resource string, parameters *gin.Params, ctx *gin.Context) (result interface{}, err types.Error)
}
