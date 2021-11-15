package main

import "fmt"

func sort2_3(a, b, c int) (int, int, int) {
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

func sort3(a, b, c int) (int, int, int) {
	return sort2_3(a, b, c)
}

func main() {
	fmt.Print("Enter 3 numbers: ")
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	a, b, c = sort3(a, b, c)
	fmt.Println("Sorted:", a, b, c)
}
