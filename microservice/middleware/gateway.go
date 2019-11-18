package middleware

import (
	"encoding/base64"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"heurd.com/wand-go/wand/abstract"
	"heurd.com/wand-go/wand/types/gateway"
)

type GatewayMiddleware struct {
	abstract.Middleware
}

func (this *GatewayMiddleware) AfterRoute (ctx *gin.Context) {

	data := gateway.Data{}

	if decodeData, ok := base64.URLEncoding.DecodeString(ctx.Request.Header.Get("X-Gateway-Data")); ok == nil {
		json.Unmarshal(decodeData, &data)
	}

	ctx.Set("GATEWAY_DATA", data)
}
