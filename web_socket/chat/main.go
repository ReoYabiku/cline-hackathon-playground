package main

import (
	"log"
	"net/http"
)

func main() {
	// ここでは、一つのプロセスに一つのhubを作成することになっている。
	// 本来は、一つのROOMに対して一つのhubを作成することになる
	// ROOM単位でクライアントを単位するのであれば、ソースコード上のhubが消滅してRedis Pub/Subでブロードキャストを実現することになる
	hub := newHub()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		createConnWithClient(w, r, hub)
	})

	log.Println("serving...")
	err := http.ListenAndServe(":5555", nil)
	if err != nil {
		log.Println(err)
	}
}
