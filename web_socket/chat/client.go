package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	hub       *Hub
	conn      *websocket.Conn
	messageCh chan []byte
}

func newClient(hub *Hub, conn *websocket.Conn) *Client {
	return &Client{
		hub:       hub,
		conn:      conn,
		messageCh: make(chan []byte, 1024),
	}
}

func (c *Client) pumpToHub() {
	for {
		_, r, err := c.conn.NextReader()
		if err != nil {
			log.Println(err)
			continue
		}

		b, err := io.ReadAll(r)
		if err != nil {
			log.Println(err)
			continue
		}

		var message Message
		err = json.Unmarshal(b, &message)
		if err != nil {
			log.Println(err)
			continue
		}

		c.hub.bloadcast(&message)
	}
}

func (c *Client) pumpFromHub() {
	// messageChにメッセージが来たらWebクライアントに送信する
	for message := range c.messageCh {
		fmt.Println("message", message)
		w, err := c.conn.NextWriter(websocket.TextMessage)
		if err != nil {
			log.Println(err)
			continue
		}

		fmt.Println("writing...")
		w.Write(message)
		w.Close()
	}
}
