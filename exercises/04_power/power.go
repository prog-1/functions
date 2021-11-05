package main

import "fmt"

func Pow(x, y uint) uint {
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
	fmt.Print("Enter x and y: ")
	var x, y uint
	fmt.Scan(&x, &y)
	fmt.Println(x, y)
	pow := Pow(x, y)
	fmt.Println("x ^ y =", pow)
} //input: 4 3 | 0 5 | 7 0 | 16 2 | 12 4  |
//output : 64  | 0   | 1   | 256  | 20736 |
