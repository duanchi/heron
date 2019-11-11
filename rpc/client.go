package rpc

import (
	"fmt"
	"runtime"
)

func Call (in IN, out *OUT) error {
	pc, _, _, _ := runtime.Caller(1)

	fmt.Println(runtime.FuncForPC(pc).Name())
	return nil
}