package main

import (
	"fmt"
	"pubsub/cache"
	"time"
	//	"github.com/akyoto/cache"
)

type aClient map[string]string

//func (topics Topics) Unsub
func main() {

	client := aClient{}
	client["hi"] = "hello"

	// New cache
	c := cache.New(5 * time.Minute)

	// Put something into the cache
	c.Set("a", client, 1*time.Minute)
	// Read from the cache
	obj, _ := c.Get("a")

	// Convert the type
	fmt.Println(obj)

	client["ok"] = "yes"
	c.Set("a", client, 1*time.Minute)
	// Read from the cache
	obj, _ = c.Get("a")

	fmt.Println(obj)

	delete(client, "ok")

	c.Set("a", client, 1*time.Minute)

	obj, _ = c.Get("a")

	fmt.Println(obj)

	fmt.Println("with new topic// ==========================================")

	client["hib"] = "hellob"

	// Put something into the cache
	c.Set("b", client, 1*time.Minute)
	// Read from the cache
	obj, _ = c.Get("b")

	// Convert the type
	fmt.Println(obj)

	client["oks"] = "yes"
	c.Set("b", client, 1*time.Minute)
	// Read from the cache
	obj, _ = c.Get("b")

	fmt.Println(obj)

	delete(client, "oks")

	c.Set("b", client, 1*time.Minute)

	obj, _ = c.Get("b")

	fmt.Println(obj)

}
