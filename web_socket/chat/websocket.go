package main

import (
	"net/http"

	"github.com/gorilla/websocket"
)

// gorilla/websocketのclientを準備する
func createConnWithClient(w http.ResponseWriter, r *http.Request, hub *Hub) {
	upgrader := &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		panic(err)
	}

	// clientを作成する
	client := newClient(hub, conn)

	// clientをhubに登録する
	hub.register(client)

	// clientからメッセージを受信する後ルーチンを走らせる
	go client.pumpToHub()

	// clientにメッセージを送信する後ルーチンを走らせる
	go client.pumpFromHub()
}
