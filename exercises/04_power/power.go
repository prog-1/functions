package main

import "fmt"

func pow(a uint, b uint) uint {
	r := a
	for i := b; i > 1; i, r = i-1, r*a {
	}
	return r
}

func main() {
	fmt.Println("This program returns power of two unsigned integers")
	fmt.Print("Enter two number: ")
	var a, b uint
	fmt.Scan(&a, &b)
	fmt.Print("Pow of ", a, " and ", b, " equals ", pow(a, b))

}
