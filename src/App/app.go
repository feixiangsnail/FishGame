package App

import (

	"Sockets"
	"net/http"
)
var Client Sockets.WebSocketClientModel
func StartServer(){
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		Client = Sockets.WebSocketClientModel{}
		if !Client.Init(w, r) {
			return
		}
		go Client.OnMessage()
	})
	http.ListenAndServe(":8082", nil)


}


