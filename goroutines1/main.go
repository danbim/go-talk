package main

import (
	"fmt"
)

func Greeter(inbox chan string, outbox chan string) {
	name := <-inbox // waits for message to be available (synchronisation point)
	outbox <- fmt.Sprintf("Hello %v!", name)
}

func main() {

	gIn := make(chan string)
	gOut := make(chan string)

	go Greeter(gIn, gOut) // starts concurrent goroutine (potentially on another thread)

	gIn <- "Thilko"    // sends message to whoever is listening on the channel
	response := <-gOut //  waits for message to be available (synchronisation point)
	fmt.Println(response)

}
