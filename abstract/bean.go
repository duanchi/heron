package abstract

type Bean struct {
	name string
}

func (this *Bean) Init () {}

func (this *Bean) GetName () (name string) {
	return this.name
}

func (this *Service) SetName (name string) {
	this.name = name
}