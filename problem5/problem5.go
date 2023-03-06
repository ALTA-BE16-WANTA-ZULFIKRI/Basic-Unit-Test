package main

import (
	"fmt"
    "math"
)

func PrintLuasPermukaan(radius, tinggi float64) float64 {
	// your code here
	luasPermukaan := 2 * math.Pi * radius * (radius + tinggi)
	return luasPermukaan
}

func main() {
	var T float64 = 20
	var r float64 = 4

	fmt.Println(PrintLuasPermukaan(r, T))
}
