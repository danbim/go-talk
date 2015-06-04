package main

import (
	"fmt"
)

func Greeter(inchan chan string, outchan chan string) {
	name := <-inchan // waits for message to be available (synchronisation point)
	outchan <- fmt.Sprintf("Hello %v!", name)
}

func main() {

	gIn := make(chan string)
	gOut := make(chan string)

	go Greeter(gIn, gOut) // starts concurrent goroutine (potentially on another thread)

	gIn <- "Thilko"    // sends message to whoever is listening on the channel
	response := <-gOut //  waits for message to be available (synchronisation point)
	fmt.Println(response)

}
