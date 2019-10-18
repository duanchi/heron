package types

type RuntimeError struct {
	Message string
	ErrorCode    int
}

func (err RuntimeError) Error() string {
	return err.Message
}

func (err RuntimeError) Code() int {
	return err.ErrorCode
}