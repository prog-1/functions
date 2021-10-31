package main

import "fmt"

func main() {
	fmt.Print("Enter three numbers: ")
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	min := a
	if b < min {
		min = b
	}
	if c < min {
		min = c
	}

	max := a
	if b > max {
		max = b
	}
	if c > max {
		max = c
	}

	fmt.Println("min =", min)
	fmt.Println("max =", max)
}
