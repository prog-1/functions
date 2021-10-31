package main

import "fmt"

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

// func minmax(a, b, c int) (mn, mx int) {
func minmax(a, b, c int) (int, int) {
	return min(a, min(b, c)), max(a, max(b, c))
}

func main() {
	fmt.Print("Enter three numbers: ")
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	mn, mx := minmax(a, b, c)

	fmt.Println("min =", mn)
	fmt.Println("max =", mx)
}
