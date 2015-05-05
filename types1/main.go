package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func showRef(v Vertex) {
	fmt.Println(v)
}

func showPtr(v *Vertex) {
	fmt.Println(*v)
}

func main() {
	v := Vertex{1, 2}
	showRef(v)
	showPtr(&v)

	w := Vertex{}
	showRef(w)
	showPtr(&w)

	x := Vertex{Y: 2}
	showRef(x)
	showPtr(&x)
}
