package abstract

import "github.com/xormplus/xorm"

type Model struct {
	Bean
	engine *xorm.Engine `json:"-" xorm:"-"`
}
