package main

import "fmt"

func pow(a uint, b uint) uint {
	res := a
	for i := b; i > 1; i, res = i-1, res*a {
	}
	return res
}

func main() {
	fmt.Println("This program returns power of two unsigned integers")
	fmt.Print("Enter two number: ")
	var a, b uint
	fmt.Print("Pow of ", a, " and ", b, " equals ", pow(a, b))

}
