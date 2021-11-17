package main

func reverse(n int) int {

	r := 0
	for ; n != 0; n /= 10 {
		r = n%10 + (r * 10)
	}
	return r
}
func main() {

}
