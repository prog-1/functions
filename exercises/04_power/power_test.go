package main

import "testing"

// TestPow tests a Pow(x, y) function, which returns integer x**y.
func TestPow(t *testing.T) {
	for _, tc := range []struct {
		x, y int
		want int
	}{
		{x: 1, y: 2, want: 1},
		{x: 2, y: 2, want: 4},
		{x: 3, y: 2, want: 9},
		{x: 4, y: 2, want: 16},
		{x: 5, y: 2, want: 25},
		{x: 6, y: 2, want: 36},
		{x: 10, y: 10, want: 10000000000},
		{x: 2, y: 3, want: 8},
		{x: 2, y: 4, want: 16},
	} {
		if got := Pow(tc.x, tc.y); got != tc.want {
			t.Errorf("power(%v, %v) = %v, want = %v", tc.x, tc.y, got, tc.want)
		}
	}
}
