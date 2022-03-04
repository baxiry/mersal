package main

//type Hub struct{}

type Client map[string]bool

func (h Hub) Join(t string, c Client) {
}

func (h Hub) Delete(t, c string) {

}
