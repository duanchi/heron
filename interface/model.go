package _interface

import "github.com/xormplus/xorm"

type ModelInterface interface {
	SetEngine (name string)
	SetDb (engine *xorm.Engine)
	Options () map[string]interface{}
}