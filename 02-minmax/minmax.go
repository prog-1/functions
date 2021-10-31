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

func main() {
	fmt.Print("Enter three numbers: ")
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	mn := min(a, min(b, c))
	mx := max(a, max(a, b))

	fmt.Println("min =", mn)
	fmt.Println("max =", mx)
}
