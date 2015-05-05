package main

import "fmt"

func genAdd(howMuch int) func(int) int {
	return func(a int) int {
		return a + howMuch
	}
}

func main() {
	addOne := genAdd(1)
	addTwo := genAdd(2)
	addThree := genAdd(3)

	fmt.Println(addOne(1))
	fmt.Println(addTwo(1))
	fmt.Println(addThree(1))
}
