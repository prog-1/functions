package main

func minmax(a, b int) (mn, mx int) {
	if a > b {
		return b, a
	}
	return a, b
}
