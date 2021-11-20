package main

import "fmt"

func pow(x, y uint) uint {
	if y == 0 {
		x = 1
		return x
	}
	for c := x; y != 1; y-- {
		x = c * x
	}
	return x
}

func main() {
	fmt.Println("Enter two numbers:")
	var x, y uint
	fmt.Scanln(&x, &y)
	power := pow(x, y)
	fmt.Println("x ^ y =", power)
}
