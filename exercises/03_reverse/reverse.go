package main

import "fmt"

func reverse(x int) int {
	var r int
	for n := x; n != 0; n /= 10 {
		r = n%10 + (r * 10)
	}
	return r
}
func main() {
	fmt.Println("This program reverses number.")
	fmt.Print("Write your number:")
	var x int
	fmt.Scan(&x)
	x = reverse(x)
	fmt.Println("The reversed number is:", x)
}
