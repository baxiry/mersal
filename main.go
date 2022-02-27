package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

type Client string // will be websocket.Client

type Subscribers map[Client]bool

type Topics map[string]Subscribers

// Subscribe adds new client to topic
// if topic is not exist then Subscribe create new
func (topics Topics) Subscribe(top string, cli Client) {
	if topics[top] == nil {
		topics[top] = Subscribers{}
	}
	//subs[cli] = true
	topics[top][cli] = true
}

//func (topics Topics) Unsub
func main() {

	var topics = Topics{}
	var cli Client
	//var subs = Subscribers{}

	for i := 0; i < 10000; i++ {
		topic := "topic-" + strconv.Itoa(i)
		cli = "client-" + Client(strconv.Itoa(i))
		topics.Subscribe(topic, cli)
	}

	os.Exit(0)
	// ---------------------

	file, err := os.Create("./test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	jsondata, err := json.Marshal(topics)
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
