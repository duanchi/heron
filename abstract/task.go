package abstract

import (
	_interface "go.heurd.com/heron-go/heron/interface"
)

type Task struct {
	Bean
	_interface.TaskInterface
}

func (this *Task) OnStart () {}

func (this *Task) OnExit () {}

func (this *Task) AfterInit () {}