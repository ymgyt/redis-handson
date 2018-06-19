package client

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/ymgyt/redis-handson/chat/core"
	"github.com/ymgyt/redis-handson/chat/datastore"
	"github.com/ymgyt/redis-handson/chat/protocol"
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

func NewClient(clientID string) *Client {
	return &Client{
		ID: clientID,
	}
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

func ConnectionByID(ID string) (*websocket.Conn, error) {
	return registry.get(ID)
}

func (c *Client) SendCloseConnection() {
	c.Connection.SetWriteDeadline(time.Now().Add(writeWait))
	c.Connection.WriteMessage(websocket.CloseMessage, nil)
}

func (c *Client) SendPing() error {
	c.Connection.SetWriteDeadline(time.Now().Add(writeWait))
	return c.Connection.WriteMessage(websocket.PingMessage, []byte("ping"))
}

func (c *Client) Setup() {
	c.Connection.SetReadLimit(maxMEssageSize)
	c.Connection.SetReadDeadline(time.Now().Add(pongWait))
	c.Connection.SetPongHandler(func(string) error {
		c.Connection.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})
}

func (c *Client) Close() error {
	registry.delete(c)
	return c.Connection.Close()
}

func (c *Client) ReadRPC() (*protocol.RPC, error) {
	mt, data, err := c.Connection.ReadMessage()
	if err != nil {
		return nil, err
	}

	if mt == 0 {
		return nil, errors.New("invalid data received")
	}

	rpc := &protocol.RPC{}
	return rpc, json.Unmarshal(data, rpc)
}

func (c *Client) SendEvent(event *protocol.Event) error {
	data, err := json.Marshal(event)
	if err != nil {
		core.Logger.Error(err)
		return err
	}
	return datastore.Redis.Publish(c.ID, data)
}
