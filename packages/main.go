package main

import "fmt"
import "github.com/danbim/go-talk/packages/vertex"

// import v "github.com/danbim/go-talk/packages/vertex" // rename "vertex" in file-local scope

func main() {
	// "cannot refer to unexported name vertex.hiddenVertex"
	// s := vertex.hiddenVertex
	s := vertex.New(1, 2)
	fmt.Println(s)
}
