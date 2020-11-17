package websocket

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/duanchi/heron/types"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Handle(id string, resource string, parameters *gin.Params, ctx *gin.Context,
	handleFunction func(connection *websocket.Conn, id string, resource string, parameters *gin.Params, ctx *gin.Context) (err types.Error),
) (err error) {

	connection, websocketError := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if websocketError != nil {
		return types.RuntimeError{
			Message:   websocketError.Error(),
			ErrorCode: http.StatusInternalServerError,
		}
	}
	defer connection.Close()
	err = handleFunction(connection, id, resource, parameters, ctx)
	if err != nil {
		return
	}

	return
}
