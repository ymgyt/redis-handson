package client

import (
	"net/http"
	"time"

	websocket "github.com/howtv/gsskt_backend/pkg/dep/sources/https---github.com-gorilla-websocket"
)

const (
	writeWait      = 30 * time.Second
	pongWait       = 30 * time.Second
	maxMEssageSize = 1024 * 1024
)

var registry = newRegistry()

type Client struct {
	ID         string
	Connection *websocket.Conn
}

func (c *Client) Close() error {
	registry.delete(c)
	return c.Connection.Close()
}

func NewFromRequest(clientID string, w http.ResponseWriter, r *http.Request) *Client {
	var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return nil
	}

	client := &Client{
		ID:         clientID,
		Connection: conn,
	}

	registry.set(client)

	return client
}
