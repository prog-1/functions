package main

import "testing"

func TestMinmax(t *testing.T) {
	for _, tc := range []struct {
		a, b, c          int
		wantMin, wantMax int
	}{
		{a: 1, b: 2, c: 3, wantMin: 1, wantMax: 3},
		{a: 2, b: 1, c: 3, wantMin: 1, wantMax: 3},
		{a: 1, b: 3, c: 2, wantMin: 1, wantMax: 3},
		{a: 3, b: 1, c: 2, wantMin: 1, wantMax: 3},
		{a: 2, b: 3, c: 1, wantMin: 1, wantMax: 3},
		{a: 3, b: 2, c: 1, wantMin: 1, wantMax: 3},

		{a: 1, b: 1, c: 2, wantMin: 1, wantMax: 2},
		{a: 1, b: 2, c: 1, wantMin: 1, wantMax: 2},
		{a: 2, b: 1, c: 1, wantMin: 1, wantMax: 2},

		{a: 1, b: 1, c: 1, wantMin: 1, wantMax: 1},
	} {
		if gotMin, gotMax := minmax(tc.a, tc.b, tc.c); gotMin != tc.wantMin || gotMax != tc.wantMax {
			t.Errorf("max(%v, %v, %v) = (%v, %v), want = (%v, %v)", tc.a, tc.b, tc.c, gotMin, gotMax, tc.wantMin, tc.wantMax)
		}
	}
}
