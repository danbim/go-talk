package main

import "fmt"

var packageVar int = 1

func main() {

	var uninit uint64

	fmt.Printf("packageVar = %v\n", packageVar)
	fmt.Printf("uninit = %v\n", uninit)

	//var i = 5
	i := uint32(5)
	fmt.Printf("i[%T] = %v\n", i, i)

	a, b, c := "eins", 2, false
	fmt.Printf("a[%T] = %v\n", a, a)
	fmt.Printf("b[%T] = %v\n", b, b)
	fmt.Printf("c[%T] = %v\n", c, c)

	const truth = true
	fmt.Println("Go rules?", truth)

	// bool, string
	// int, int8, int16, int32, int64
	// uint, uint8, uint16, uint32, uint64, uintptr
	// byte (alias for uint8)
	// rune (alias for int32, represents a Unicode code point)
	// float32, float64
	// complex64, complex128
}
