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

}
func Testfunc() {
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
