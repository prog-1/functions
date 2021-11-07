package main

import "fmt"

func minmax(a, b int) (mn, mx int) {
	if a <= b {
		return a, b
	}
	return b, a
}

func sort3(a, b, c int) (int, int, int) {
	a, b = minmax(a, b)
	b, c = minmax(b, c)
	a, b = minmax(a, b)
	b, c = minmax(b, c)
	return a, b, c
}

func main() {
	fmt.Print("Enter 3 numbers: ")
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	a, b, c = sort3(a, b, c)
	fmt.Println("Sorted:", a, b, c)
}
