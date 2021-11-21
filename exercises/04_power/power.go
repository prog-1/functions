package main

import "fmt"

func Pow(x, y uint) uint {
	if y == 0 {
		x = 1
		return x
	}
	for z := x; y != 1; y-- {
		x = z * x
	}
	return x
}

func main() {
	fmt.Println("Enter two numbers:")
	var x, y uint
	fmt.Scan(&x, y)
	integer := Pow(x, y)
	fmt.Println(x, "^", y, "=", integer)
}
