package ws

import (
	"github.com/gorilla/websocket"
)

type ClientConn struct {
	UserID string
	Conn   *websocket.Conn

	Send chan []byte // writeBuffer(size = 64)
}
