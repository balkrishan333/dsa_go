package main

import (
	"fmt"
)

func main() {
	fmt.Print(longestPalindrome("babad"), "\n")
}

func longestPalindrome(s string) string {

	start, end, length := 0, 0, 0

	for i := 0; i < len(s); i++ {
		odd := expand(s, i, i)
		even := expand(s, i, i+1)

		length = max(odd, even)

		if length > end-start {
			start = i - (length-1)/2
			end = i + length/2
		}
	}
	return s[start : end+1]
}

func expand(s string, left int, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return (right - 1) - (left + 1) + 1
}
