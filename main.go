package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/tidwall/gjson"
)

type msg struct {
	Num int
}

var h Helper

func wsHandler(w http.ResponseWriter, r *http.Request) {
	//if r.Header.Get("Origin")!="http://"+r.Host {http.Error(w,"Origin not allowed",-1);return}
	fmt.Println("new client")

	conn, err := websocket.Upgrade(w, r, w.Header(), 2, 2) //1024, 1024)
	if err != nil {
		http.Error(w, "Could not open websocket connection", 404)
	}
	go echo(conn)
}

func echo(conn *websocket.Conn) {

	/* note: income data must be json as :
		{
			"event":"",
			"msg":""
		}
	event must be subscribe, unsubscriber, close or msg
	*/

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

			Publishe(i, channel, []byte(data))

		} else if event == "subscribe" {

			Subscribe(channel, conn)
			msg = []byte("subscribe to " + channel + " success!")

		} else if event == "unsubscribe" {

			Unsubscribe(channel, conn)

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

var mt sync.Mutex

func main() {
	http.HandleFunc("/ws", wsHandler)
	panic(http.ListenAndServe(":8080", nil))
}

// ---------------------------------------------------------

//http.HandleFunc("/", rootHandler)
/*
func rootHandler(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadFile("index.html")
	if err != nil {
		fmt.Println("Could not open file.", err)
	}
	fmt.Fprintf(w, "%s", content)
}
*/
