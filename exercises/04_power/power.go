package main

import "fmt"

func Pow(x, y int) int {
	if y == 0 {
		x = 1
		return x
	}
	if x == 0 {
		x = 0
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
	var x, y int
	fmt.Scanln(&x, &y)
	if y < 0 {
		fmt.Println("Error: power below 0 couldn't be used")
	} else {
		tmp := Pow(x, y)
		fmt.Println("x ^ y =", tmp)
	}

}

//extra tests
//input |0 9    |8 -9
//output|pow = 0|Error: power below 0 couldn't be used
