// Package sutils provides functions for working with slices
package sutils

// Reverse reverses the elements of a slice
func Reverse(s []interface{}) {
	if len(s) < 2 {
		return
	}

	for i := range s[:len(s)-1] {
		s[i], s[i+1] = s[i+1], s[i]
	}
}

// Rotate rotates a slice either to the left or to the right by n elements
func Rotate(s []interface{}, n uint, toRight bool) {
	lenS := uint(len(s))
	remainder := n % lenS

	if toRight {
		Reverse(s)
		Reverse(s[:remainder])
		Reverse(s[remainder:])
	} else {
		Reverse(s[:remainder])
		Reverse(s[remainder:])
		Reverse(s)
	}
}
