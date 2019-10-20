package abstract

import (
	"github.com/gin-gonic/gin"
	"heurd.com/wand-go/wand/types"
)

type Controller struct {
}

func (controller *Controller) Fetch (id string, resource string, parameters *gin.Params, ctx *gin.Context) (result interface{}, err types.Error) {
	return "error", nil
}

func (controller *Controller) Create (id string, resource string, parameters *gin.Params, ctx *gin.Context) (result interface{}, err types.Error) {
	return "error", nil
}

func (controller *Controller) Update (id string, resource string, parameters *gin.Params, ctx *gin.Context) (result interface{}, err types.Error) {
	return "error", nil
}

func (controller *Controller) Remove (id string, resource string, parameters *gin.Params, ctx *gin.Context) (result interface{}, err types.Error) {
	return "error", nil
}