package abstract

import (
	_interface "github.com/duanchi/heron/interface"
)

type Task struct {
	Bean
	_interface.TaskInterface
}

func (this *Task) OnStart () {}

func (this *Task) OnExit () {}

func (this *Task) AfterInit () {}