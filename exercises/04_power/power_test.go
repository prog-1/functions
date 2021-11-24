package main

import "testing"

// TestPow tests a Pow(x, y) function, which returns integer x**y.
func TestPow(t *testing.T) {
	for _, tc := range []struct {
		x, y int
		want int
	}{
		{1, 1, 1},
		{121, 0, 1},
		{2, 4, 16},
		{3, 5, 243},
		{10, 2, 100},
		{11, 2, 121},
		{-5, 3, -125},
		{-4, 4, 256},
		{0, 8, 0},
	} {
		if got := Pow(tc.x, tc.y); got != tc.want {
			t.Errorf("Pow(%v, %v) = %v, want = %v", tc.x, tc.y, got, tc.want)
		}
	}
}
