package main

import (
	"fmt"
	"pubsub/pubsub"

	"github.com/gorilla/websocket"
	"github.com/tidwall/gjson"
)

//note: income data must be json as `{"event":"","msg":""}`
// event must be subscribe, unsubscriber, close or msg
func serveMessages(conn *websocket.Conn) {

	for {

		i, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message no.", i)
			conn.Close()
			return
		}

		// un/subscribe if event == un/subscribe.
		event := gjson.Get(string(msg), "event").String()
		channel := gjson.Get(string(msg), "channel").String()
		data := gjson.Get(string(msg), "data").String()

		if event == "message" {

			pubsub.Publishe(i, channel, []byte(data))

		} else if event == "subscribe" {

			pubsub.Subscribe(channel, conn)
			msg = []byte("subscribe to " + channel + " success!")

		} else if event == "unsubscribe" {

			pubsub.Unsubscribe(channel, conn)
			msg = []byte("unsubscribe from " + channel + " success!")
		}

		fmt.Printf(string(msg))

		mt.Lock()
		if err = conn.WriteMessage(i, []byte("done")); err != nil {
			fmt.Println(err)
		}
		mt.Unlock()
	}
}
