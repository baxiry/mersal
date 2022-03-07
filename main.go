package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/tidwall/gjson"
)

type msg struct {
	Num int
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Origin") != "http://"+r.Host {
		//http.Error(w, "Origin not allowed", 403)
		//return
	}
	conn, err := websocket.Upgrade(w, r, w.Header(), 512, 512) //1024, 1024)
	if err != nil {
		http.Error(w, "Could not open websocket connection", 404)
	}
	go echo(conn)
}

func echo(conn *websocket.Conn) {

	// TODO use pubsub here

	for {

		fmt.Printf("%s\n", conn.RemoteAddr().String())
		i, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message no.", i)
			conn.Close()
			return
		}
		// subscribe if event == subscribe
		jmsg := gjson.Get(string(msg), "event")
		if jmsg.String() == "subscribe" {

			fmt.Println("new subscriber")

			//Subscribe("channel", "adam")
			continue
		}

		fmt.Printf("Got message: %#v\n", string(msg))

		if err = conn.WriteMessage(i, msg); err != nil {
			fmt.Println(err)
		}
	}
}

func main() {
	http.HandleFunc("/ws", wsHandler)

	panic(http.ListenAndServe(":8080", nil))
}

// ---------------------------------------------------------

//http.HandleFunc("/", rootHandler)
func rootHandler(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadFile("index.html")
	if err != nil {
		fmt.Println("Could not open file.", err)
	}
	fmt.Fprintf(w, "%s", content)
}
