package main

import "testing"

// TestPow tests a Pow(x, y) function, which returns integer x**y.
func TestPow(t *testing.T) {
	for _, tc := range []struct {
		x, y uint
		want uint
	}{
		{x: 0, y: 0, want: 1},
		{x: 1, y: 1, want: 1},
		{x: 0, y: 7, want: 0},
		{x: 7, y: 0, want: 1},
		{x: 1, y: 5, want: 1},
		{x: 5, y: 1, want: 5},
		{x: 7, y: 3, want: 343},
		{x: 8, y: 4, want: 4096},
		{x: 567, y: 4, want: 103355177121},
		{x: 12, y: 15, want: 15407021574586368},
	} {
		if got := Pow(tc.x, tc.y); got != tc.want {
			t.Errorf("pow(%v, %v) = %v, want = %v", tc.x, tc.y, got, tc.want)
		}
	}
}
