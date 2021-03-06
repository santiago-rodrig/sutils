package sutils

import "testing"

func TestReverseInts(t *testing.T) {
	s := []int{1, 2, 3}
	want := []int{3, 2, 1}
	ReverseInts(s)

	if !equalInts(want, s) {
		t.Errorf("want %v, got %v", want, s)
	}
}

func equalInts(s1, s2 []int) bool {
	if len(s1) != len(s1) {
		return false
	}

	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}
