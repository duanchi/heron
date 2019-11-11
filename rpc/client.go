package rpc

import "runtime"

func Call (in IN, out *OUT) error {
	pc, _, _, _ := runtime.Caller(1)

	runtime.FuncForPC(pc).Name()
	return nil
}