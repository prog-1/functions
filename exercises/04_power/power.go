package main

import "fmt"

func p(x, y uint) uint {
	if y == 0 {
		x = 1
		return x
	}

	for tmp := x; y != 1; y-- {
		x = x * tmp
	}
	return x
}
func main() {
	fmt.Println("Program determines x to the power of y")
	fmt.Println("Enter x and y:")
	var x, y uint
	fmt.Scanln(&x, &y)
	power := p(x, y)
	fmt.Println("x ^ y =", power)
}
