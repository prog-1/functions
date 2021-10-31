package main

import "testing"

func TestMin(t *testing.T) {
	if got := min(1, 2); got != 1 {
		t.Errorf("min(1, 2) = %v, want = 1", got)
	}
	if got := min(4, 3); got != 3 {
		t.Errorf("min(4, 3) = %v, want = 3", got)
	}
	if got := min(5, 5); got != 5 {
		t.Errorf("min(5, 5) = %v, want = 5", got)
	}
}

func TestMax(t *testing.T) {
	for _, tc := range []struct {
		a, b int
		want int
	}{
		{a: 1, b: 2, want: 2},
		{a: 4, b: 3, want: 4},
		{a: 5, b: 5, want: 5},
	} {
		if got := max(tc.a, tc.b); got != tc.want {
			t.Errorf("max(%v, %v) = %v, want = %v", tc.a, tc.b, got, tc.want)
		}
	}
}
