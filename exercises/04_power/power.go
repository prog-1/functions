package main

import "fmt"

func pow(a int, b int) int {
	p := 1
	for ; b > 0; b-- {
		p *= a
	}
	return p
}

func main() {
	fmt.Println("This program returns power of two unsigned integers")
	fmt.Print("Enter two number: ")
	var a, b int
	fmt.Print("Pow of ", a, " and ", b, " equals ", pow(a, b))

}
