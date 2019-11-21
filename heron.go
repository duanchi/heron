package heron

import (
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
	"go.heurd.com/heron-go/heron/bean"
)

var HttpServer *gin.Engine
var Db *xorm.Engine
var Config interface{}

func GetBean(name string) interface{} {
	return bean.Get(name)
}


