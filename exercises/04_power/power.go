package main

import "fmt"

func Pow(x, y uint) (uint, uint) {
	if y == 0 {
		x = 1
	} else {
		for p := x; y != 1; y-- {
			x = x * p
		}
	}

	return x, y
}
func main() {
	fmt.Print("Enter numbers: ")
	var x, y uint
	fmt.Scan(&x, &y)
	x, y = Pow(x, y)
	fmt.Println(x, y)
}
