package main

import (
	"fmt"
)

func Greeter(inbox chan string, outbox chan string, gQuit chan int, bla string) {
	for {
		select {
		case name := <-inbox:
			outbox <- fmt.Sprintf("Hello %v %v!", bla, name)
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

	go Greeter(gIn, gOut, gQuit, "No1")
	go Greeter(gIn, gOut, gQuit, "No2")

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
