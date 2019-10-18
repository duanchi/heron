package util

import (
	"os"
	"reflect"
)

func GetType (i interface{}) reflect.Type {
	return reflect.TypeOf(i).Elem()
}

func Getenv (key string, defaults string) string {
	result := os.Getenv(key)

	if result == "" {
		return defaults
	} else {
		return result
	}
}