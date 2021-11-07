package main

import "fmt"

func reverse(n int) int {
	y := 0
	for x := n; x != 0; x /= 10 {
		y = y*10 + x%10
	}
	return y
}

func main() {
	fmt.Println("Enter the number:")
	var n int
	fmt.Scan(&n)
	n = reverse(n)
	fmt.Println("Reversed:", n)
}
