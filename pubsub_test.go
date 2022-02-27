package main

import (
	"testing"
)

type args struct {
	arg1, arg2, result string
}

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
