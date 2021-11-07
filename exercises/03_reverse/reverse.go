package main

func reverse(n int) int {
	result := 0
	for ; n != 0; n /= 10 {
		result = (result*10 + (n % 10))
	}
	return result
}

func main() {
}
