package main

import "fmt"

func reverse(n int) int {
	var s = 0
	for ; n != 0; n /= 10 {
		s = s*10 + (n % 10)
	}
	return s
}

func main() {
	fmt.Print("Enter number: ")
	var n int
	fmt.Scan(&n)
	n = reverse(n)
	fmt.Println("Reversed:", n)
}
