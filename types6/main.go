package main

import "fmt"

type Stringer interface {
	String() string
}

type Vertex struct {
	X, Y float64
}

type Line struct {
	Start Vertex
	End   Vertex
}

func (v Vertex) String() string {
	return fmt.Sprintf("Vertex{X=%v, Y=%v}}", v.X, v.Y)
}

func (l Line) String() string {
	return fmt.Sprintf("Line{Start=%v, End=%v}", l.Start, l.End)
}

func main() {
	var v Stringer
	v = Vertex{1, 2}
	// v.X = 3 // does not work, v is of type Stringer
	fmt.Println(v.String())

	v = Line{Vertex{1, 2}, Vertex{3, 4}}
	fmt.Println(v.String())
}
