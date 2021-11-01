package main

func reverse(n int) int {
	result := 0
	for ; n != 0; n /= 10 {
		result = (result + (n % 10)) * 10
	}
	return result / 10
}

func main() {
}
