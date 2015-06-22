package main

import "fmt"
import "github.com/danbim/go-talk/packages2/vertex"
import "github.com/danbim/go-talk/packages2/line"

func main() {
	var s vertex.Vertex
	s = vertex.New(1, 2)
	s.MoveByLine(line.Line{vertex.New(2, 3), vertex.New(3, 2)})
	fmt.Println(s)
}
