package main

import (
	"fmt"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	var f MyFloat
	f = MyFloat(-13.0)
	fmt.Println(f)
}
