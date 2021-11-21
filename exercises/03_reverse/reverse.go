package main

import "fmt"

func Reverse(n int) (r int) {
	for ; n != 0; n /= 10 {
		r = r*10 + n%10
	}
	return r
}

func main() {
	fmt.Println("Enter the number you want to reverse:")
	var n int
	fmt.Scan(&n)
	n = Reverse(n)
	fmt.Println("The reversed number is", n)
}
