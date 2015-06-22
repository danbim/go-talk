package main

import (
	"fmt"
	"math/rand"
	"time"
)

func doQuery(service string) string {
	// simulate some hard work or I/O
	time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
	return service
}

func main() {

	for i := 0; i < 10; i++ {

		ch := make(chan string, 3)
		go func() { ch <- doQuery("https://www.google.com/") }()
		go func() { ch <- doQuery("https://www.bing.com/") }()
		go func() { ch <- doQuery("https://www.yahoo.com/") }()

		now := time.Now()
		select {
		case resp := <-ch:
			fmt.Printf("Got response from %v after %v\n", resp, time.Since(now))
		case <-time.After(100 * time.Millisecond):
			fmt.Printf("Timed out :(\n")
		}
	}
}
