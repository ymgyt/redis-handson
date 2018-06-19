package events

import (
	"github.com/ymgyt/redis-handson/chat/messaging/client"
	"github.com/ymgyt/redis-handson/chat/protocol"
)

type Event struct {
	*protocol.Event
}

func NewEvent(t protocol.Event_Type) *Event {
	pe := &protocol.Event{
		Type: t,
	}

	return &Event{pe}
}

func (e *Event) SendToClient(clientID string) {
	c := client.NewClient(clientID)
	c.SendEvent(e.Event)
}
