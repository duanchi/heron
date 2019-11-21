package abstract

import _interface "go.heurd.com/heron-go/heron/interface"

type Service struct {
	_interface.ServiceInterface
	Bean
}

func (this *Service) Init () {
	this.Bean.Init()
}

func (this *Service) GetServiceName () (name string) {
	return this.BeanName
}

func (this *Service) SetServiceName (name string) {
	this.BeanName = name
}