package messaging

import (
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	websocket "github.com/howtv/gsskt_backend/pkg/dep/sources/https---github.com-gorilla-websocket"
	"github.com/ymgyt/redis-handson/chat/datastore"
)

const (
	writeWait  = 30 * time.Second
	pingPeriod = 10 * time.Second
)

func Start(r *mux.Router) {
	go publishListener()
	r.Methods("GET").Path("/{client_id}").HandlerFunc(handler)
}

func publishListener() {
	datastore.Redis.Subscribe(func(channel string, data []byte) {
		chunks := strings.Split(channel, ":")
		clientID := chunks[len(chunks)-1]
		conn, err := client.ConnectionByID(clientID)
		if err != nil {
			core.Logger.Errorf("MESSAGING: Subscribe error %s %v %v", clientID, string(data), err)
			return
		}
		core.Logger.Info("MESSAGING: Subscribe %s %v", clientID, string(data))
		conn.SetWriteDeadline(time.Now().Add(writeWait))
		conn.WriteMessage(websocket.TextMessage, data)
	})
}

func handler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	client := client.NewFromRequest(params["client_id"], w, r)
	if client == nil {
		core.Logger.Error("Unauthorized")
		return
	}

	ch := make(chan *protocol.RPC)
	go dispatcher(client, ch)
	reader(client, ch)
	close(ch)
}
