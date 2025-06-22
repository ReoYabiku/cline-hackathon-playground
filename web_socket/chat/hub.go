package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
)

type Hub struct {
	clients *sync.Map
}

func newHub() *Hub {
	return &Hub{
		clients: &sync.Map{},
	}
}

func (h *Hub) register(client *Client) error {
	// 既存のclientには何もしない
	if _, ok := h.clients.Load(client); ok {
		return nil
	}

	// 新規のクライアントの場合、hubに登録する
	h.clients.Store(client, struct{}{})
	return nil
}

func (h *Hub) bloadcast(message *Message) error {
	fmt.Println("bloadcast...")
	h.clients.Range(func(key, _ any) bool {
		fmt.Println(1)
		client, ok := key.(*Client)
		if !ok {
			// keyがclientじゃないのでスキップ
			return true
		}

		fmt.Println(2)

		b, err := json.Marshal(message)
		if err != nil {
			log.Println(err)
			return true
		}

		fmt.Println(3)

		client.messageCh <- b
		return true
	})
	return nil
}

func (h *Hub) unregister(client *Client) error {
	h.clients.Delete(client)
	return nil
}
