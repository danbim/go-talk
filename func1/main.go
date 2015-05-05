package main

import "fmt"

// "normal" definition
func add(x int, y int) int {
	return x + y
}

// short form
func add2(x, y int) int {
	return x + y
}

// multiple return values
func swap(x, y int) (int, int) {
	return y, x
}

func main() {
	if add(1, 2) != add2(swap(1, 2)) {
		panic("D'oh!")
	}
	fmt.Println("¡Ay, caramba!")
}
