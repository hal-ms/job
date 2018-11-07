package socket

import (
	"github.com/googollee/go-socket.io"
)

var Server *socketio.Server

func init() {
	server, err := socketio.NewServer(nil)
	if err != nil {
		panic(err)
	}
	server.On("connection", func(so socketio.Socket) {
		so.Join("all")
	})
	Server = server
}

func SendAll(event string, args interface{}) {
	Server.BroadcastTo("all", event, args)
}
