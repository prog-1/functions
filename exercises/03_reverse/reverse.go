package main

import "fmt"

func reverse(a int) (y int) {
	for tmp := a; tmp != 0; tmp /= 10 {
		y = y*10 + tmp%10
	}
	return y
}
func main() {
	fmt.Println("This program reverse a number.")
	fmt.Print("Enter a number: ")
	var a int
	fmt.Scan(&a)
	a = reverse(a)
	fmt.Println("reverse is ", a)

}
