package main

import "fmt"

func g() {
	// do something sensible, then panic
	panic("Aargh!")
}

func h() {
	r := recover()
	if r != nil {
		fmt.Printf("Recovered from panic (%v) in f()\n", r)
	}
}

func f() {
	defer h()
	fmt.Println("Entering f(), calling g()")
	g()
	fmt.Println("Returned normally from g()")
}

func main() {
	f() // causes program to crash
}
