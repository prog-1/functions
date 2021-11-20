package main

import (
	"fmt"
)

func reverse(n int) int {
	var y int
	for z := n; z != 0; z /= 10 {
		y = y*10 + z%10
	}
	return y
}

func main() {
	fmt.Print("Enter the number:")
	var n int
	fmt.Scan(&n)
	n = reverse(n)
	fmt.Println("The reversed number is:", n)
}
