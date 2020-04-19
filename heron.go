package heron

import (
	"github.com/gin-gonic/gin"
	"github.com/xormplus/xorm"
	"go.heurd.com/heron-go/heron/bean"
	"go.heurd.com/heron-go/heron/log"
)

var HttpServer *gin.Engine
var Db *xorm.Engine
var Config interface{}
var Log *log.Logrus
var Model []*xorm.Engine

func GetBean(name string) interface{} {
	return bean.Get(name)
}


