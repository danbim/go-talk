package main

import (
	"fmt"
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

	gIn := make(chan string)
	gOut := make(chan string)
	gQuit := make(chan int)

	go Greeter(gIn, gOut, gQuit)

	gIn <- "Thilko"
	fmt.Println(<-gOut)

	gIn <- "Sven"
	fmt.Println(<-gOut)

	gIn <- "Tobias"
	fmt.Println(<-gOut)

	gIn <- "Sina"
	fmt.Println(<-gOut)

	gQuit <- 1

}
