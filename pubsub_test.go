package main

import (
	"fmt"
	"strconv"
	"testing"
)

// test Topics.Subscribe() method
func TestSubscribe(t *testing.T) {

}

// test CreateTopic
func TestCreateTopic(t *testing.T) {

	// some test cases with expected result
	cases := []struct {
		arg1, arg2, result string
	}{
		{"123", "456", "415263"},
		{"723", "356", "735263"},
		// add more test cases here
	}

	for _, c := range cases {
		if CreateTopic(c.arg1, c.arg2) != c.result {
			t.Errorf("CreateTopic func error: %s with %s must return %s resutl \n", c.arg1, c.arg2, c.result)

		}
	}
}

func BenchmarkSubscribe(b *testing.B) {
	hub := newHub()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < b.N/2; i++ {
			topic := "topic-" + strconv.Itoa(i)
			cli := "client-" + Client(strconv.Itoa(i))
			hub.Subscribe(topic, cli)
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := b.N / 2; i < b.N; i++ {
			topic := "topic-" + strconv.Itoa(i)
			cli := "client-" + Client(strconv.Itoa(i))
			hub.Subscribe(topic, cli)
		}
	}()

	wg.Wait()
	fmt.Println(len(hub.Topics))
	lendata := 0
	for _, v := range hub.Topics {
		lendata += len(v)
	}
	fmt.Println(len(hub.Topics))

}
