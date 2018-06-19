package client

import (
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

type clientRegistry struct {
	mutex       sync.RWMutex
	connections map[string]*websocket.Conn
}

func newRegistry() *clientRegistry {
	return &clientRegistry{connections: map[string]*websocket.Conn{}}
}

func (c *clientRegistry) set(client *Client) {
	c.mutex.Lock()
	c.connections[client.ID] = client.Connection
	c.mutex.Unlock()
}

func (c *clientRegistry) get(id string) (*websocket.Conn, error) {
	c.mutex.RLock()
	conn, ok := c.connections[id]
	c.mutex.RUnlock()

	if !ok {
		return nil, fmt.Errorf("could not find client with client ID %s", id)
	}
	return conn, nil
}

func (c *clientRegistry) delete(client *Client) {
	c.mutex.Lock()
	delete(c.connections, client.ID)
	c.mutex.Unlock()
}
