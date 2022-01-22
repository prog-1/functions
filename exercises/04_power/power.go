package main

import "fmt"

func Pow(x, y uint) uint {
	if y == 0 {
		x = 1
		return x // because all situation, when y = 0 -> x = 1. Test without it doesn't work (by x^y)
	}
	for z := x; y != 1; y-- {
		x = z * x
	}
	return x
}

func main() {
	fmt.Println("Enter x and y")
	var x, y uint
	fmt.Scan(&x, &y)
	power := Pow(x, y)
	fmt.Println(x, "^", y, "=", power)
}
