package abstract

type Service struct {
	Bean
	name string
}

func (this *Service) Init () {
	this.Bean.Init()
}

func (this *Service) GetServiceName () (name string) {
	return this.name
}

func (this *Service) SetServiceName (name string) {
	this.name = name
}