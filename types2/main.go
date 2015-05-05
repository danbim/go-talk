package main

import "fmt"

func main() {

	a := []int{0, 1, 2, 3, 4}
	printSlice("a", a)

	// creates a new sub-slices [low:high)
	a1 := a[:2]
	a2 := a[2:5]

	printSlice("a1", a1)
	printSlice("a2", a2)

	b := make([]int, 5)    // type, length
	c := make([]int, 0, 5) // type, length, capacity

	printSlice("b", b)
	printSlice("c", c)

	a = append(a, 3) // increases length, doubles capacity
	printSlice("a", a)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
