package main

import (
	// "fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

var (
	upgrader = websocket.Upgrader{
		// 允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func wsHandler(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Hello"))
	var (
		wsConn *websocket.Conn
		err    error
		conn   *impl.Connection
	)

	// upgrade: websocket
	if wsConn, err = upgrader.Upgrade(w, r, nil); err != nil {
		return
	}
	if conn, err = impl.InitConnection(wsConn); err != nil {
		goto ERR
	}

	go func() {
		var (
			err error
		)
		for {

			if err = conn.WriteMessage([]byte("Heartbeat")); err != nil {
				return
			}
			time.Sleep(1 * time.Second)
		}
	}()

	for {
		if data, err = conn.ReadMessage(); err != nil {
			goto ERR
		}
		if err = conn.WriteMessage(data); err != nil {
			goto ERR
		}
	}

	// upgrade: websocket
	// if conn, err = upgrader.Upgrade(w, r, nil); err != nil {
	// 	return
	// }
	// for {
	// 	// Text binary
	// 	if _, data, err = conn.ReadMessage(); err != nil {
	// 		goto ERR
	// 	}
	// 	if err = conn.WriteMessage(websocket.TextMessage, data); err != nil {
	// 		goto ERR
	// 	}

	// }

ERR:
	conn.Close()
	// TODO: Close connection
}

func main() {
	http.HandleFunc("/ws", wsHandler)
	http.ListenAndServe("0.0.0.0:8888", nil)
}
