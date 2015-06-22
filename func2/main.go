package main

import "fmt"
import "strings"

func bla(sentence string, modifier func(string) string) {
	fmt.Println(modifier(sentence))
}

func main() {

	say := bla

	say("¡Ay, caramba!", strings.ToUpper)
}
