package im

import (
	"log"

	"github.com/gorilla/websocket"
	"github.com/tidwall/gjson"
)

// note: income data must be json as `{"event":"","msg":""}`
// event must be join, leave, close or msg
func ServeMessages(conn *websocket.Conn) {

	for {

		i, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message no.", i)
			conn.Close()
			continue
		}

		// handle messages based on event.
		var smsg = string(msg)

		event := gjson.Get(smsg, "event").String()
		// TODO continue if no event.

		channel := gjson.Get(smsg, "channel").String()

		switch event {
		case "msg": //
			Publish(i, channel, []byte(gjson.Get(smsg, "data").String()))

		case "join":
			Subscribe(channel, conn)
			msg = []byte("join to " + channel)

			if err = conn.WriteMessage(i, msg); err != nil {
				log.Println("no event", err)
			}

		case "leave":
			Unsubscribe(channel, conn)
			msg = []byte("left from " + channel)

			if err = conn.WriteMessage(i, msg); err != nil {
				log.Println("no event", err)
			}

		default:

			if err = conn.WriteMessage(i, []byte("unknown envet. please read docs")); err != nil {
				log.Println("no event", err)
			}
		}

		/*
			if event == "message" {
				Publish(i, channel, []byte(data))

			} else if event == "subscribe" {
				Subscribe(channel, conn)
				msg = []byte("subscribe to " + channel + " success!")

			} else if event == "unsubscribe" {
				Unsubscribe(channel, conn)
				msg = []byte("unsubscribe from " + channel + " success!")
			}

			//fmt.Println(string(msg))

			err = conn.WriteMessage(i, msg)

			if err != nil {
				log.Println(err)
			}
		*/
	}

}
