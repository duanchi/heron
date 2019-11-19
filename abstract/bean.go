package abstract

import (
	_interface "go.heurd.com/heron-go/heron/interface"
)

type Bean struct {
	_interface.BeanInterface
	BeanName string
}

func (this *Bean) Init () {}

func (this *Bean) GetBeanName () (name string) {
	return this.BeanName
}

func (this *Bean) SetBeanName (name string) {
	this.BeanName = name
}