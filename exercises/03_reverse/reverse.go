package main

import "fmt"

func reverse(n int) int {
	var y int
	for tmp := n; tmp != 0; tmp /= 10 {
		y = y*10 + tmp%10
	}
	return y
}

func main() {
	fmt.Print("Enter number: ")
	var n int
	fmt.Scan(&n)
	n = reverse(n)
	fmt.Println("Reversed number is:", n)
} //input: 821 | 019 | 17 | -12 | 100 |
//output : 128 | 1   | 71 | -21 | 1   |
