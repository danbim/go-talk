package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

func main() {

	var m = make(map[string]Vertex)                    // "normal" instantiation
	m["nbsp Hackerspace"] = Vertex{53.86826, 10.68554} // set
	fmt.Println(m["nbsp Hackerspace"])                 // get

	var n = map[string]Vertex{ // struct literal
		"Bell Labs": Vertex{40.68433, -74.39967},
		"Google":    {37.42202, -122.08408}, // struct literal with inferred type
	}

	fmt.Println(n)
}
