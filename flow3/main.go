package main

import "fmt"

func f() {
	defer fmt.Println("Leaving f()") // pushed on stack, executed on leaving f(), even when panicking
	fmt.Println("Entering f()")
	panic("Aargh!")
}

func main() {
	f() // causes program to crash
}
