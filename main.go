package main

import (
	"fmt"
	"pubsub/cache"
	"strconv"
	//	"github.com/akyoto/cache"
)

var c = cache.New()

func Subscribe(topic, client string) {
	clients, _ := c.Get(topic)
	if clients == nil {
		clients = make(map[string]bool)
	}
	clients.(map[string]bool)[client] = true

	c.Set(topic, clients)
}

func Unsubscribe(topic, client string) {

	clients, _ := c.Get(topic)
	if clients == nil {
		return
	}

	delete(clients.(map[string]bool), client)
	c.Set(topic, clients)

}

func Publishe(t string) {
	clients, _ := c.Get(t)
	for k := range clients.(map[string]bool) {

		fmt.Println("    data sent to ", k)
	}
}

func main() {
	// test Subscribe
	for i := 0; i < 10; i++ {
		topic := "topic-" + strconv.Itoa(i)
		for i := 0; i < 100; i++ {
			client := "client-" + strconv.Itoa(i)
			Subscribe(topic, client)

		}

	}

	for i := 0; i < 10; i++ {
		topic := "topic-" + strconv.Itoa(i)
		fmt.Println("start Publishe to all Subscriber in", topic)

		Publishe(topic)
	}

}
