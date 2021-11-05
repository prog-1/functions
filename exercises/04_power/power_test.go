package main

import "testing"

// TestPow tests a Pow(x, y) function, which returns integer x**y.
func TestPow(t *testing.T) {
	for _, tc := range []struct {
		x, y uint
		want uint
	}{
		{1, 1, 1},
		{2, 3, 8},
		{5, 2, 25},
		{13, 1, 13},
		{0, 2, 0},
		{3, 0, 1},
		{7, 0, 1},
		{3, 2, 9},
		{2, 2, 4},
		{4, 3, 64},
	} {
		if got := Pow(tc.x, tc.y); got != tc.want {
			t.Errorf("Pow(%v, %v) = (%v), want(%v)",
				tc.x, tc.y, got, tc.want)
		}
	}
}
