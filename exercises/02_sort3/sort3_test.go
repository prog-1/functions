package main

import "testing"

func TestSort4(t *testing.T) {
	for _, tc := range []struct {
		a, b, c             int
		wantA, wantB, wantC int
	}{
		{1, 2, 3, 1, 2, 3},
		{1, 3, 2, 1, 2, 3},
		{2, 1, 3, 1, 2, 3},
		{2, 3, 1, 1, 2, 3},
		{3, 1, 2, 1, 2, 3},
		{3, 2, 1, 1, 2, 3},
	} {
		gotA, gotB, gotC := sort3(tc.a, tc.b, tc.c)
		if gotA != tc.wantA || gotB != tc.wantB || gotC != tc.wantC {
			t.Errorf("sort4(%v, %v, %v) = (%v, %v, %v), want = (%v, %v, %v)",
				tc.a, tc.b, tc.c, gotA, gotB, gotC, tc.wantA, tc.wantB, tc.wantC)
		}
	}
}
