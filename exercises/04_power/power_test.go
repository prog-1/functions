package main

import "testing"

// TestPow tests a Pow(x, y) function, which returns integer x**y.
func TestPow(t *testing.T) {
	for _, tc := range []struct {
		x, y uint
		want uint
	}{
		{x: 0, y: 0, want: 1},
		{x: 0, y: 1, want: 0},
		{x: 1, y: 0, want: 1},
		{x: 2, y: 3, want: 8},
		{x: 2, y: 9, want: 512},
		{x: 3, y: 3, want: 27},
	} {
		if gotx, goty := Pow(tc.x, tc.y); gotx != tc.want && goty >= 0 {
			t.Errorf("Pow(%v, %v) = (%v), want = (%v)", tc.x, tc.y, gotx, tc.want)
		}
	}
}
