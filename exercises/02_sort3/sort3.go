package main

import "fmt"

func sort3(a, b, c int) (int, int, int) { // 3 4 5
	if b < a {
		a, b = b, a
	}
	if c < a {
		a, c = c, a
	}
	if c < b {
		b, c = c, b
	}

	return a, b, c
}

func main() {
	fmt.Print("Enter 3 numbers: ")
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	a, b, c = sort3(a, b, c)
	fmt.Println("Sorted:", a, b, c)
}
