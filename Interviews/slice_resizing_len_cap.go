package main

import "fmt"

func SliceResizingLenCap() {
	// Playing with slice headers (len, cap, and pointer).
	// 1. All slices created from the same underlying array share the same memory until you reslice beyond capacity or append.
	s := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(s)
	a := s[:0]
	fmt.Println(a)
	a = append(a, 90)
	// a = append(a, 90)
	// a = append(a, 90)
	// a = append(a, 90)
	// a = append(a, 90)
	// a = append(a, 90)

	a = a[:3]

	fmt.Println(a) // [90, 2,3]

	// IF 3 append(a, 90) and a = a[:6] THEN output looks like [90, 90, 90, 4, 5, 6]
}
