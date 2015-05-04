package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.

func fibN(n int) int {
	if n==0||n==1 {
		return n
	} else {
		return fibN(n-1) + fibN(n-2)
	}
}

func fibonacci() func(int) int {
	return fibN
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}
