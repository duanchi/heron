package abstract

import (
	"github.com/gin-gonic/gin"
	"heurd.com/wand-go/wand/types"
)

type RestController struct {
	Bean
}

func (THIS *RestController) Fetch (id string, resource string, parameters *gin.Params, ctx *gin.Context) (result interface{}, err types.Error) {
	return "error", nil
}

func (THIS *RestController) Create (id string, resource string, parameters *gin.Params, ctx *gin.Context) (result interface{}, err types.Error) {
	return "error", nil
}

func (THIS *RestController) Update (id string, resource string, parameters *gin.Params, ctx *gin.Context) (result interface{}, err types.Error) {
	return "error", nil
}

func (THIS *RestController) Remove (id string, resource string, parameters *gin.Params, ctx *gin.Context) (result interface{}, err types.Error) {
	return "error", nil
}