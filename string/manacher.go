package main

import (
	"strings"
)

func findLongestPalindromicSubstring(s string) string {
	t := preProcess(s)
	n := len(t)
	p := make([]int, n) // p[i] stores the length of the palindrome centered at t[i]

	center, rightBoundary := 0, 0 // center is the center of the current palindrome, rightBoundary is its right edge
	maxLen, maxCenter := 0, 0

	for i := 1; i < n-1; i++ { // Iterate from the first '#' to the last '#'
		// iMirror is the mirror of i with respect to the current center
		iMirror := 2*center - i

		// If i is within the current palindrome's right boundary,
		// we can leverage the information from its mirror.
		if rightBoundary > i {
			p[i] = min(rightBoundary-i, p[iMirror])
		} else {
			p[i] = 0 // Otherwise, start with a palindrome length of 0 (just the character itself)
		}

		// Attempt to expand the palindrome centered at i.
		// The sentinels '^' and '$' prevent out-of-bounds access.
		for t[i+1+p[i]] == t[i-1-p[i]] {
			p[i]++
		}

		// If the palindrome centered at i expands beyond the current rightBoundary,
		// update the center and rightBoundary.
		if i+p[i] > rightBoundary {
			center = i
			rightBoundary = i + p[i]
		}

		// Update the longest palindrome found so far.
		if p[i] > maxLen {
			maxLen = p[i]
			maxCenter = i
		}
	}

	// Extract the longest palindromic substring from the original string s.
	// (maxCenter - 1 - maxLen) / 2 gives the starting index in the original string.
	// maxLen is the length of the palindrome in the original string.
	start := (maxCenter - 1 - maxLen) / 2
	return s[start : start+maxLen]
}

// Preprocess the string to handle even length palindromes.
//
// For example, "aba" becomes "^#a#b#a#$"
// For "abba" becomes "^#a#b#b#a#$"
func preProcess(s string) string {
	var sb strings.Builder
	sb.WriteRune('^') // Start sentinel
	for _, char := range s {
		sb.WriteRune('#')
		sb.WriteRune(char)
	}
	sb.WriteString("#$") // End sentinel
	return sb.String()
}
