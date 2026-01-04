package ws

import (
	"sync"
)

type DocRoom struct {
	DocID     string
	Clients    map[string]*ClientConn
	Register   chan *ClientConn
	Unregister chan *ClientConn
	Broadcast  chan []byte
	Mu         sync.RWMutex
}

func NewDocRoom(DocID string) *DocRoom{
	room := &DocRoom{
		DocID: DocID,
		Clients: make(map[string]*ClientConn),
		Register: make(chan *ClientConn),
		Unregister: make(chan *ClientConn),
		Broadcast: make(chan []byte),
	}

	go room.run()
	return room
}

func (r *DocRoom) run(){
	for{
		select{
		case client := <-r.Register:
			r.Mu.Lock()
			r.Clients[client.UserID] = client
			r.Mu.Unlock()

		case client := <-r.Unregister:
			r.Mu.Lock()
			if _, ok := r.Clients[client.UserID]; ok{
				delete(r.Clients, client.UserID)
				close(client.Send)
			}
			r.Mu.Unlock()

		case msg := <- r.Broadcast:
			r.Mu.RLock()
			for _, client := range r.Clients{
				select {
				case client.Send <-msg:

				default:
					go func(c *ClientConn){
							r.Unregister <- c
					}(client)
				}
			}
			r.Mu.RUnlock()
		}
	}
}
