package main

import "fmt"

func PrintLuas(alas, tinggi float64) float64 {
	// your code here
	luas := 0.5 * alas * tinggi
	return luas
}

func main() {
	var alas float64 = 20
	var tinggi float64 = 25

	fmt.Println(PrintLuas(alas, tinggi))
}
