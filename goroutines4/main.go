package main

import (
	"fmt"
	"math/rand"
	"time"
)

func doQuery(service string, ch chan string) {
	// simulate some hard work or I/O
	time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
	ch <- "Awesome event!"
}

func main() {

	for i := 0; i < 10; i++ {

		ch := make(chan string)
		go doQuery("https://www.softwerkskammer.org/groups/luebeck", ch)

		select {
		case resp := <-ch:
			fmt.Printf("Go resp=%v\n", resp)
		case <-time.After(100 * time.Millisecond):
			fmt.Printf("Timed out :(\n")
		}
	}
}
