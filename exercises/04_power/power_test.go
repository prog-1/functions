package main

import "testing"

// TestPow tests a Pow(x, y) function, which returns integer x**y.
func TestPow(t *testing.T) {
	for _, tc := range []struct {
		x, y uint
		want uint
	}{
		{x: 0, y: 0, want: 1},
		{x: 1, y: 0, want: 1},
		{x: 23, y: 0, want: 1},
		{x: 0, y: 1, want: 0},
		{x: 0, y: 176, want: 0},
		{x: 1, y: 1, want: 1},
		{x: 2, y: 3, want: 8},
		{x: 3, y: 2, want: 9},
		{x: 4, y: 8, want: 65_536},
		{x: 9, y: 6, want: 531_441},
		{x: 5, y: 5, want: 3125},
		{x: 7, y: 7, want: 823_543},
	} {
		if got := Pow(tc.x, tc.y); got != tc.want {
			t.Errorf("Pow(%v, %v) = (%v), want = (%v)", tc.x, tc.y, got, tc.want)
		}
	}
}
