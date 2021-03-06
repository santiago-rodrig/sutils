// Package sutils provides functions for working with slices
package sutils

// ReverseInts reverses the elements of a slice
func ReverseInts(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// RotateInts rotates a slice either to the left or to the right by n elements
func RotateInts(s []int, n uint, toRight bool) {
	lenS := uint(len(s))
	remainder := n % lenS

	if toRight {
		ReverseInts(s)
		ReverseInts(s[:remainder])
		ReverseInts(s[remainder:])
	} else {
		ReverseInts(s[:remainder])
		ReverseInts(s[remainder:])
		ReverseInts(s)
	}
}
