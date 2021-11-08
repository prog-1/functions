package main

import "fmt"

func Pow(x, y uint) uint {
	if y == 0 {
		x = 1
		return x
	}
	for i := x; y != 1; y-- {
		x = i * x

	}
	return x

}

func main() {
	fmt.Println("Enter the value of x and y")
	var x, y uint
	fmt.Scan(&x, &y)
	k := Pow(x, y)
	fmt.Println("x ^ y = ", x, "^", y, "=", k)
}
