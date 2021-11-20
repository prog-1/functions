package main

import (
	"fmt"
)

func pow(x, y uint) uint {
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
	fmt.Print("Enter two numbers: ")
	var x, y uint
	fmt.Scan(&x, &y)
	integer := pow(x, y)
	fmt.Println(x, "^", y, "=", integer)
}
