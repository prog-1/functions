package main

import "fmt"

func reverse(n int) int {
	var x int
	for y := n; y != 0; y /= 10 {
		x = x*10 + y%10
	}
	return x
}

func main() {
	fmt.Println("Enter the number you want to reverse:")
	var n int
	fmt.Scan(&n)
	n = reverse(n)
	fmt.Println("The reversed number is", n)
}
