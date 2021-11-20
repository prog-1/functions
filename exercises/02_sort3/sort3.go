package main

import "fmt"

func minmax(a, b int) (min, max int) {
	if a < b {
		return a, b
	}
	return b, a
}

func sort3(a, b, c int) (int, int, int) {
	a, b = minmax(a, b)
	b, c = minmax(b, c)
	a, b = minmax(a, b)
	return a, b, c
}

func main() {
	fmt.Print("Enter 3 numbers: ")
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	a, b, c = sort3(a, b, c)
	fmt.Println("Sorted numbers:", a, b, c)
} //input: 3 1 2 | 2 -3 0 | 654 389475 364 | 0 1 0 | number |
//output : 1 2 3 | -3 0 2 | 364 654 389475 | 0 0 1 | 0 0 0  |
