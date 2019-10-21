package wand

import (
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
)

var HttpServer *gin.Engine
var Db *xorm.Engine
var Config interface{}
var Bean interface{}

