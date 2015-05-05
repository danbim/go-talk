package main

import "fmt"
import "strings"

func main() {

	say := func(sentence string, modifier func(string) string) {
		fmt.Println(modifier(sentence))
	}

	say("¡Ay, caramba!", strings.ToUpper)
}
