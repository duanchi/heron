package wand

import (
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
	"heurd.com/wand-go/wand/bean"
)

var HttpServer *gin.Engine
var Db *xorm.Engine
var Config interface{}

func GetBean(name string) interface{} {
	return bean.Get(name)
}


