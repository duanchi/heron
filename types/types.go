package types

import (
	"github.com/gin-gonic/gin"
)

type Route struct {
	Url	string
	Method string
	Handler HandleFunc
}

type HandleFunc func(ctx *gin.Context)

type Response struct {
	RequestId string `json:"request_id"`
	Status bool `json:"status"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

type Error interface {
	error
	Code() int
}