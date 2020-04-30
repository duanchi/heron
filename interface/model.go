package _interface

type ModelInterface interface {
	Options () map[string]interface{}
	Table() string
	Source() string
}