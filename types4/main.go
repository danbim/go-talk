package main

import "fmt"

type Vertex struct {
	X, Y float64
}

// no classes in Go, but "methods" with "receivers"
func (v Vertex) ToString() string {
	return fmt.Sprintf("Vertex{X=%v, Y=%v}}", v.X, v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X *= f
	v.Y *= f
}

func main() {
	v := Vertex{3.0, 4.0}
	v.Scale(2.0)                    // call-by-reference
	fmt.Println("v=", v.ToString()) // call-by-value (creates a copy of the struct)
}
