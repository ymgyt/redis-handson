package rpc

import (
	"github.com/ymgyt/redis-handson/chat/messaging/client"
	"github.com/ymgyt/redis-handson/chat/messaging/events"
	"github.com/ymgyt/redis-handson/chat/protocol"
)

type typing struct{}

func newTyping() *typing {
	return &typing{}
}

func (t *typing) Start(c *client.Client, requestID string, params *protocol.RPC_TypingStartParams) {
	e := events.NewEvent(protocol.Event_TYPING_START)
	e.Event.ResponseTo = c.ID
	e.SendToClient(requestID)
}

func (t *typing) End(c *client.Client, requestID string, params *protocol.RPC_TypingEndParams) {
	e := events.NewEvent(protocol.Event_TYPING_END)
	e.Event.ResponseTo = c.ID
	e.SendToClient(requestID)
}
