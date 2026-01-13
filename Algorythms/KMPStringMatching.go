package main

import "fmt"

func KMPStringMatching(text, pattern string) bool {
	n := len(text)
	m := len(pattern)
	if m == 0 {
		return true // empty pattern is always found
	}
	if n < m {
		return false
	}

	lps := make([]int, len(pattern)) // LPS: Longest Prefix Suffix
	length, i := 0, 1

	for i < m {
		if pattern[i] == pattern[length] {
			length++
			lps[i] = length
			i++
		} else {
			if length != 0 {
				length = lps[length-1] // fall back to the length of the longest proper prefix which is also a suffix
			} else {
				lps[i] = 0 // no need but just to make it explicit
				i++
			}
		}
	}

	i, j := 0, 0

	for i < n {
		if text[i] == pattern[j] {
			j++
			i++
		}
		if j == m {
			return true
		}

		if i < n && text[i] != pattern[j] {
			if j != 0 {
				j = lps[j-1]
			} else {
				i++
			}
		}
	}

	return false
}

func RunKMP() {
	text := "ABABDABACDABABCABAB"
	pattern := "ABABCABAB"

	fmt.Println(KMPStringMatching(text, pattern)) // true
	fmt.Println(KMPStringMatching(text, "XYZ"))   // false
}
