package main

import "testing"

// TestPow tests a Pow(x, y) function, which returns integer x**y.
func TestPow(t *testing.T) {
	for _, tc := range []struct {
		x, y uint
		want uint
	}{
		// Test cases.
		{x: 0, y: 0, want: 1},
		{x: 1, y: 1, want: 1},
		{x: 0, y: 1, want: 0},
		{x: 1, y: 0, want: 1},
		{x: 0, y: 4, want: 0},
		{x: 5, y: 0, want: 1},
		{x: 2, y: 1, want: 2},
		{x: 1, y: 4, want: 1},
		{x: 3, y: 4, want: 81},
		{x: 5, y: 2, want: 25},
		{x: 12, y: 3, want: 1728},
	} {
		// Test body.
		if got := Pow(tc.x, tc.y); got != tc.want {
			t.Errorf("pow(%v, %v) = %v, want = %v", tc.x, tc.y, got, tc.want)
		}
	}
}
