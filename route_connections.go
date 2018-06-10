package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"fmt"
)

type Client struct {
	Conn *websocket.Conn
}

func newConnectionHandler(appState *AppState) gin.HandlerFunc {
	wsUpgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	return func(ctx *gin.Context) {
		userId := ctx.Param("user_id")
		conn, err := wsUpgrader.Upgrade(ctx.Writer, ctx.Request, nil)

		if err != nil {
			fmt.Println(err)
			log(fmt.Sprintf("Fail to establish WebSocket connection with %s", userId))
			return
		}

		client := Client{
			Conn: conn,
		}
		appState.Clients[userId] = client

		sendInitialServiceStates(userId, appState)
	}
}
