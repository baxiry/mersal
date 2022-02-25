package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

type Msg struct {
	From string
	To   string
	Msg  string
}

type Client string

type Topics map[string]Subscribers

type Subscribers map[Client]bool

func (topics Topics) Subscribe(top string, subs Subscribers, cli Client) {
	subs[cli] = true
	topics[top] = subs
}

func main() {

	os.Exit(0)

	var topics = Topics{}
	var subs = Subscribers{}
	var cli Client

	for i := 0; i < 1000; i++ {
		topic := "topic-" + strconv.Itoa(i)
		cli = "client-" + Client(strconv.Itoa(i))
		topics.Subscribe(topic, subs, cli)
	}

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
