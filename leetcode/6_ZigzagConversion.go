package main

import "strings"

func convert(s string, numRows int) string {
	if len(s) < 2 || numRows == 1 {
		return s
	}

	var result strings.Builder
	for row := 0; row < numRows; row++ {
		for i := row; i < len(s); i += 2 * (numRows - 1) {
			result.WriteString(string(s[i]))
			if row != 0 && row != numRows-1 && i+2*(numRows-1)-2*row < len(s) {
				result.WriteString(string(s[i+2*(numRows-1)-2*row]))
			}
		}
	}
	return result.String()
}

func main() {
	println(convert("PAYPALISHIRING", 3))
}
