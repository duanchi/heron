package abstract

import (
	_interface "go.heurd.com/heron-go/heron/interface"
)

type Bean struct {
	_interface.BeanInterface
	BeanName string
}

func (this *Bean) Init () {}

func (this *Bean) GetName () (name string) {
	return this.BeanName
}

func (this *Bean) SetName (name string) {
	this.BeanName = name
}