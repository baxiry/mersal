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

	/*
		for topic, subs := range topics {
			fmt.Println()
			fmt.Print("\n", topic, "\t")

			for k, v := range subs {
				fmt.Print(" ", k, v)
			}
		}
	*/

	jsondata, err := json.Marshal(topics)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}
	err = ioutil.WriteFile("temp.txt", []byte(jsondata), 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(" Done")

	//pull := map[string]bool{}
	//msg := Msg{"ahmed", "d7ome", "hello my frend"}
}
