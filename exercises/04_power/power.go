package main

import "fmt"

func Pow(x, y float64) float64 {
	num := x
	if y < 0 {
		for ; y < -1; y++ {
			num = num * x
		}
		num = 1 / num
	}
	if y == 0 {
		num = 1
	}
	for ; y > 1; y-- {
		num = num * x
	}
	return num
}

func main() {
	fmt.Print("Enter the number and the power:")
	var x, y float64
	fmt.Scan(&x, &y)
	num := Pow(x, y)
	fmt.Println(num)
}
