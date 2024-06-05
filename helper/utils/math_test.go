package utils

import "testing"

func TestM_Min(t *testing.T) {
	min := Math.Min(6, 1, 34, 5, 3)
	if min != 1 {
		t.Fatal("min error")
	}
	max := Math.Max(8, 1, 34, 5)
	if max != 34 {
		t.Fatal("min error")
	}
}
