package main

import "fmt"

func reverse(n int) int {
	var y int
	for ; n != 0; n /= 10 {
		y = y*10 + n%10
	}
	return y
}

func main() {
	fmt.Println("Enter the number")
	var n int
	fmt.Scan(&n)
	n = reverse(n)
	fmt.Println("The reversed number is", n)
}
