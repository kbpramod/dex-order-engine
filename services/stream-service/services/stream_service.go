package services

import (
	"log"

	"github.com/gorilla/websocket"
)

func StartStream() {

	url := "wss://stream.binance.com:9443/ws/btcusdt@ticker"

	conn, _, err := websocket.DefaultDialer.Dial(url, nil)

	if err != nil {
		log.Fatal("connection error:", err)
	}

	defer conn.Close()

	log.Println("Connected to Binance")

	for {
		_, message, err := conn.ReadMessage()

		if err != nil {
			log.Println("read error:", err)
			return
		}

		log.Println(string(message))
	}
}