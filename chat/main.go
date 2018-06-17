package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	messaging.Start(r)
	http.ListenAndServe(":3000", r)
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
