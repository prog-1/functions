package main

import "fmt"

func reverse(num int) (rev int) {
	for tmp := 0; num > 0; num /= 10 {
		tmp = num % 10
		rev = rev*10 + tmp
	}
	return rev
}

func main() {
	fmt.Print("Enter the number:")
	var num int
	fmt.Scan(&num)
	num = reverse(num)
	fmt.Println("Reversed:", num)
}
