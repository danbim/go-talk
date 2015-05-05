package main

import "fmt"

func g() {
	// do something sensible, then panic
	panic("Aargh!")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic (%v) in f()\n", r)
		}
	}()
	fmt.Println("Entering f(), calling g()")
	g()
	fmt.Println("Returned normally from g()")
}

func main() {
	f() // causes program to crash
}
