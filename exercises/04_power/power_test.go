package main

import "testing"

// TestPow tests a Pow(x, y) function, which returns integer x**y.
func TestPow(t *testing.T) {
	for _, tc := range []struct {
		x, y uint
		want uint
	}{
		{x: 0, y: 0, want: 1},
		{x: 2, y: 0, want: 1},
		{x: 0, y: 9, want: 0},
		{x: 14, y: 9, want: 20661046784},
		{x: 6, y: 9, want: 10077696},
		{x: 0, y: 234, want: 0},
		{x: 78, y: 0, want: 1},
		{x: 2, y: 2, want: 4},
		{x: 5, y: 5, want: 3125},
		{x: 4, y: 5, want: 1024},
	} {
		if got := pow(tc.x, tc.y); got != tc.want {
			t.Errorf("pow(%v, %v) = %v, want = %v", tc.x, tc.y, got, tc.want)
		}
	}
}
