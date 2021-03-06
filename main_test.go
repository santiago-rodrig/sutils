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

func TestRotateInts(t *testing.T) {
	t.Run("to the left 2 places", func(t *testing.T) {
		s := []int{1, 2, 3}
		want := []int{3, 1, 2}
		RotateInts(s, 2, false)

		if !equalInts(want, s) {
			t.Errorf("want %v, got %v", want, s)
		}
	})

	t.Run("to the right 2 places", func(t *testing.T) {
		s := []int{1, 2, 3}
		want := []int{2, 3, 1}
		RotateInts(s, 2, true)

		if !equalInts(want, s) {
			t.Errorf("want %v, got %v", want, s)
		}
	})
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
