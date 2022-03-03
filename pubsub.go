package main

type Client map[string]bool

type Topics map[string]bool

func (topics Topics) Join(t string, c Client) {
	topics[t] = c
}

func (topics Topics) Delete(t, c string) {

	delete(topics[t], c)
}
