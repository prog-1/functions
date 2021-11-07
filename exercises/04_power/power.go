package main

import "fmt"

func pow(a int, b int) int {
	if b == 0 {
		return 1
	}
	res := a
	for i := b; i > 1; i, res = i-1, res*a {
	}
	return res
}

func main() {
	fmt.Println("This program returns power of two unsigned integers")
	fmt.Print("Enter two number: ")
	var a, b int
	fmt.Print("Pow of ", a, " and ", b, " equals ", pow(a, b))

}
