package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	
	d := make([][]uint8, dy)
	for y, _ := range d {
		
		d[y] = make([]uint8, dx)
		for x, _ := range d[y] {
			d[y][x] = uint8((x+y)/2)
		}
	}
	
	return d
}

func main() {
	pic.Show(Pic)
}
