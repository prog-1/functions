package main

import "fmt"

func reverse(number int) int {
	y := 0
	for x := number; x != 0; x /= 10 {
		y = y*10 + x%10
	}
	return y
}

func main() {
	fmt.Println("Enter the number:")
	var num int
	fmt.Scanln(&num)
	num = reverse(num)
	fmt.Println("reserve:", num)
}
