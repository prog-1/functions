package main

import "fmt"

func Pow(x, y int) int {
	r := x

	for i := y; i > 1; i-- {
		r *= x
	}
	return r

}
func main() {
	fmt.Println("This program determines power of x and y.")
	fmt.Print("Write x and y:")
	var x, y int
	fmt.Scan(&x, &y)
	r := Pow(x, y)
	fmt.Printf("Power of %v and %v equals %v.", x, y, r)
}
