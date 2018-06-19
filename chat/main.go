package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ymgyt/redis-handson/chat/messaging"
)

const (
	host = ":4000"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	messaging.Start(r)
	fmt.Println("listening ...", host)
	panic(http.ListenAndServe(host, r))
}

/*
var socket = new WebSocket("ws://localhost:3000/user__1");
socket.onopen= function() {
  console.log('ok');
}

socket.addEventListener('message', function(e) {
  console.log(JSON.parse(e.data))
})


*/
