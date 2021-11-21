package main

import "testing"

// TestPow tests a Pow(x, y) function, which returns integer x**y.
func TestPow(t *testing.T) {
	for _, tc := range []struct {
		x, y uint
		want uint
	}{
		{x: 1, y: 2, want: 1},
		{x: 2, y: 10, want: 1024},
		{x: 3, y: 5, want: 243},
		{x: 4, y: 2, want: 16},
		{x: 5, y: 2, want: 25},
		{x: 13, y: 3, want: 2197},
		{x: 10, y: 10, want: 10000000000},
		{x: 1000, y: 0, want: 1},
		{x: 0, y: 0, want: 1},
	} {
		if got := Pow(tc.x, tc.y); got != tc.want {
			t.Errorf("power(%v, %v) = %v, want = %v", tc.x, tc.y, got, tc.want)
		}
	}
}
