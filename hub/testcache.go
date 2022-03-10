package main

import "fmt"

func mainc() {
	c := NewCache()
	c.Set("hi0", "hello")
	c.Set("hi1", "hello1")
	c.Set("hi2", "hello2")
	c.Set("hi3", "hello3")

	hi0, _ := c.Get("hi0")
	hi1, _ := c.Get("hi1")
	hi2, _ := c.Get("hi2")
	hi3, _ := c.Get("hi3")
	fmt.Println(hi0)
	fmt.Println(hi1)
	fmt.Println(hi2)
	fmt.Println(hi3)
}
