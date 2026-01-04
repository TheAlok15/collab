package ws

func (c *ClientConn) readloop(r *DocRoom) {

	defer func() {

		r.Unregister <- c

		c.Conn.Close()
	}()

	for {
		_,msg,err := c.Conn.ReadMessage()
		if err != nil{
			return
		}

		r.Broadcast <- msg
	}

}