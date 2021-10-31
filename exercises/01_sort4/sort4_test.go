package main

import "testing"

func TestSort4(t *testing.T) {
	for _, tc := range []struct {
		a, b, c, d                 int
		wantA, wantB, wantC, wantD int
	}{
		{1, 1, 1, 1, 1, 1, 1, 1},
		{4, 1, 1, 1, 1, 1, 1, 4},
		{4, 4, 1, 1, 1, 1, 4, 4},
		// https://play.golang.org/p/frIM7BG90rn
		{1, 2, 3, 4, 1, 2, 3, 4},
		{1, 2, 4, 3, 1, 2, 3, 4},
		{1, 3, 2, 4, 1, 2, 3, 4},
		{1, 3, 4, 2, 1, 2, 3, 4},
		{1, 4, 3, 2, 1, 2, 3, 4},
		{1, 4, 2, 3, 1, 2, 3, 4},
		{2, 1, 3, 4, 1, 2, 3, 4},
		{2, 1, 4, 3, 1, 2, 3, 4},
		{2, 3, 1, 4, 1, 2, 3, 4},
		{2, 3, 4, 1, 1, 2, 3, 4},
		{2, 4, 3, 1, 1, 2, 3, 4},
		{2, 4, 1, 3, 1, 2, 3, 4},
		{3, 2, 1, 4, 1, 2, 3, 4},
		{3, 2, 4, 1, 1, 2, 3, 4},
		{3, 1, 2, 4, 1, 2, 3, 4},
		{3, 1, 4, 2, 1, 2, 3, 4},
		{3, 4, 1, 2, 1, 2, 3, 4},
		{3, 4, 2, 1, 1, 2, 3, 4},
		{4, 2, 3, 1, 1, 2, 3, 4},
		{4, 2, 1, 3, 1, 2, 3, 4},
		{4, 3, 2, 1, 1, 2, 3, 4},
		{4, 3, 1, 2, 1, 2, 3, 4},
		{4, 1, 3, 2, 1, 2, 3, 4},
		{4, 1, 2, 3, 1, 2, 3, 4},
	} {
		gotA, gotB, gotC, gotD := sort4(tc.a, tc.b, tc.c, tc.d)
		if gotA != tc.wantA || gotB != tc.wantB || gotC != tc.wantC || gotD != tc.wantD {
			t.Errorf("sort4(%v, %v, %v, %v) = (%v, %v, %v, %v), want = (%v, %v, %v, %v)",
				tc.a, tc.b, tc.c, tc.d, gotA, gotB, gotC, gotD, tc.wantA, tc.wantB, tc.wantC, tc.wantD)
		}
	}
}
