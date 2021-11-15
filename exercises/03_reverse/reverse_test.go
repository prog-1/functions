package main

import "testing"

func TestReverse(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want int
	}{
		{n: 0, want: 0},
		{n: 1, want: 1},
		{n: 11, want: 11},
		{n: 12, want: 21},
		{n: 123, want: 321},
		{n: 1223, want: 3221},
		{n: 10_000, want: 1},
		{n: -234, want: -432},
	} {
		if got := reverse(tc.n); got != tc.want {
			t.Errorf("reverse(%v) = %v, want = %v", tc.n, got, tc.want)
		}
	}
}
