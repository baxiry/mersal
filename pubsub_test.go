package main

import (
	"testing"
)

// args that CreateTopin need with resutl that it must return
type args struct {
	arg1, arg2, result string
}

// some test cases with expected result
var cases = []args{
	{"123", "456", "415263"},
	{"723", "356", "735263"},

	// add more test cases here
}

// test CreateTopic
func TestCreateTopic(t *testing.T) {

	for _, c := range cases {
		if CreateTopic(c.arg1, c.arg2) != c.result {
			t.Errorf("CreateTopic func error: %s with %s must return %s resutl \n", c.arg1, c.arg2, c.result)

		}
	}
}
