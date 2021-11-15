package main

import "fmt"

func reverse(a int) int {
	var b int
	for c := a; c != 0; c /= 10 {
		b = b*10 + c%10
	}
	return b
}

func main() {
	fmt.Println("Enter the number:")
	var a int
	fmt.Scanln(&a)
	a = reverse(a)
	fmt.Println("The reversed number is:", a)
}
