package main

import (
	"fmt"
	//"time"
)

func Greeter(inbox chan string, outbox chan string, gQuit chan int) {
	for {
		var name string
		select {
		case name = <-inbox:
			outbox <- fmt.Sprintf("Hello %v!", name)
		case <-gQuit:
			fmt.Println("Done for the day :)")
			return
		}
	}
}

func main() {

	gIn := make(chan string, 1)
	gOut := make(chan string, 1)
	gQuit := make(chan int)

	go Greeter(gIn, gOut, gQuit)

	gIn <- "Thilko"
	gIn <- "Sven"
	gIn <- "Tobias"
	//	gIn <- "Sina" // will produce deadlock :(

	fmt.Println(<-gOut)
	fmt.Println(<-gOut)
	fmt.Println(<-gOut)
	//	fmt.Println(<-gOut)

	gQuit <- 1
}
