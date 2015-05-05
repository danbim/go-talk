package main

import "github.com/danbim/go-talk/mytypes"
import "fmt"

func main() {
	// "cannot refer to unexported name mytypes.internalString"
	// var s mytypes.internalString
	var s mytypes.ExportedString
	s = "hello world"
	fmt.Println(s)
}
