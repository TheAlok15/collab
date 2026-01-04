package ws

import "github.com/gorilla/websocket"

func (c *ClientConn) writeLoop(r *DocRoom) {

	defer func() {

		r.Unregister <- c
		c.Conn.Close()

	}()

	for {
		for msg := range c.Send {
			err := c.Conn.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				return
			}

		}
	}
}