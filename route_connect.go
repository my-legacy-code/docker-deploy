package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"fmt"
	"log"
	"net/http"
)

type Client struct {
	Conn *websocket.Conn
}

func newConnectionHandler(appState *AppState, errLogger *log.Logger) gin.HandlerFunc {
	wsUpgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	return func(ctx *gin.Context) {
		userId := ctx.Param("user_id")
		conn, err := wsUpgrader.Upgrade(ctx.Writer, ctx.Request, nil)

		if err != nil {
			fmt.Println(err)
			errLogger.Printf("Fail to establish WebSocket connection with %s", userId)
			return
		}

		client := Client{
			Conn: conn,
		}
		appState.Clients[userId] = client

		sendInitialServiceStates(userId, appState)
	}
}
