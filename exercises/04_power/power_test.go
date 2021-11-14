package main

import "testing"

// TestPow tests a Pow(x, y) function, which returns integer x**y.
func TestPow(t *testing.T) {
	for _, tc := range []struct {
		x, y float64
		want float64
	}{
		{x: 0, y: 0, want: 1},
		{x: 8, y: 0, want: 1},
		{x: 10, y: -1, want: 0.1},
		{x: 5, y: -3, want: 0.008},
		{x: -3, y: 3, want: -27},
		{x: -4, y: 5, want: -1024},
		{x: -2, y: 4, want: 16},
		{x: -7, y: 2, want: 49},
		{x: 6, y: 2, want: 36},
		{x: 1, y: 4, want: 1},
		{x: 2, y: 3, want: 8},
	} {
		if got := Pow(tc.x, tc.y); got != tc.want {
			t.Errorf("Pow(%v, %v) = %v, want = %v", tc.x, tc.y, got, tc.want)
		}
	}
}
