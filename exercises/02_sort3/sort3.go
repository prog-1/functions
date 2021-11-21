package main

import "fmt"

func minmax(a, b int) (mn, mx int) {
	if b < a {
		return b, a
	}
	return a, b
}

func Sort3(a, b, c int) (int, int, int) {
	a, b = minmax(a, b)
	b, c = minmax(b, c)
	a, b = minmax(a, b)
	return a, b, c
}

func main() {
	fmt.Print("Enter 3 numbers: ")
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	a, b, c = Sort3(a, b, c)
	fmt.Println("Sorted:", a, b, c)
}
