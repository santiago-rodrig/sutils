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

func TestRemoveAdjacentDuplicates(t *testing.T) {
	s := []string{"hello", "you", "you", "you", "hey", "hey", "car"}
	want := []string{"hello", "you", "hey", "car"}
	s = RemoveAdjacentDuplicates(s)

	if !equalStrings(s, want) {
		t.Errorf("want %v, got %v", want, s)
	}
}

func TestSquashSpaces(t *testing.T) {
	s := "Hello   there\n\n\tmy friend"
	want := "Hello there\nmy friend"
	got := SquashSpaces([]byte(s))

	if string(got) != want {
		t.Errorf("want %q, got %q", want, string(got))
	}
}

func TestReverseUTF8Bytes(t *testing.T) {
	s := "Hello there"
	want := "ereht olleH"
	got := string(ReverseUTF8Bytes([]byte(s)))

	if got != want {
		t.Errorf("want %v, got %v", want, got)
	}
}

func equalStrings(s1, s2 []string) bool {
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
