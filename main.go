package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type Client string // will be websocket.Client

//type Subscribers map[Client]bool

type Hub struct {
	//mt     sync.Mutex
	Subscribers map[Client]bool
	Topics      map[string]map[Client]bool
}

// Subscribe adds new client to topic
// if topic is not exist then Subscribe create new
func (h Hub) Subscribe(top string, cli Client) {

	h.Subscribers[cli] = true
	h.Topics[top] = h.Subscribers
}
func newHub() *Hub {
	return &Hub{
		Subscribers: make(map[Client]bool),
		Topics:      make(map[string]map[Client]bool, 1),
	}
}

//func (topics Topics) Unsub
func main() {
	hub := newHub()

	for i := 0; i < 5; i++ {
		topic := "topic-" + strconv.Itoa(i)
		cli := "client-" + Client(strconv.Itoa(i))
		hub.Subscribe(topic, cli)
	}

	fmt.Println(len(hub.Topics["topic-1"]))

	for _, t := range hub.Topics {
		fmt.Println(t)
		for _, v := range t {
			fmt.Println("   ", v)
		}
	}

	os.Exit(0)
	// ---------------------

	jsondata, err := json.Marshal(hub)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}

	fmt.Println(" Done")

	Updatefile("temp.txt", jsondata)

	//pull := map[string]bool{}
	//msg := Msg{"ahmed", "d7ome", "hello my frend"}
}

// createTopic create new topic id topic from clients ids
func CreateTopic(id1, id2 string) (topic string) {

	for i := 0; i < len(id1); i++ {
		if id1[i] > id2[i] {
			topic += string(id1[i]) + string(id2[i])
		} else {
			topic += string(id2[i]) + string(id1[i])
		}
	}
	return topic
}

// update file updates jsondata condtent file
func Updatefile(filePath string, jsondata []byte) error {
	err := ioutil.WriteFile("temp.txt", jsondata, 0644)
	if err != nil {
		return err
	}
	return nil
}

// Msg message type for testing
type Msg struct {
	From string
	To   string
	Msg  string
}
