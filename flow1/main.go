package main

import "fmt"

func main() {

	sum := 0
	for i := 0; i < 10; i++ { // no brackets for condition, but mandatory curly brackets
		sum += i
	}
	fmt.Println(sum)

	sum2 := 1
	for sum2 < 1000 { // for = while
		sum2 += sum2
	}
	fmt.Println(sum2)

	switch sum2 { // no breaks necessary
	case 1024:
		fmt.Println("case 1024!")
	case 768:
		fmt.Println("case 768!")
	default:
		fmt.Println("Neither 1024 nor 768")
	}
}
