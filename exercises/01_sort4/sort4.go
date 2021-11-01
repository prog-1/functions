package main

import "fmt"

func minmax(a, b int) (mn, mx int) {
	return min(a, b), max(a, b)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func sort4(a, b, c, d int) (int, int, int, int) {
	a, b = minmax(a, b)
	b, c = minmax(b, c)
	c, d = minmax(c, d)
	a, b = minmax(a, b)
	b, c = minmax(b, c)
	a, b = minmax(a, b)
	return a, b, c, d
}

func main() {
	fmt.Print("Enter 4 numbers: ")
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	a, b, c, d = sort4(a, b, c, d)

	fmt.Println("Sorted numbers:", a, b, c, d)
}
