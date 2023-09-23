package main

func lengthOfLongestSubstring(s string) int {

	if len(s) == 1 {
		return 1
	}

	indexes := make([]int, 255)

	for i := 0; i < len(indexes); i++ {
		indexes[i] = -1
	}

	length, start := 0, 0

	for i := 0; i < len(s); i++ {
		var ch byte = s[i]
		var index int = indexes[ch]

		if index != -1 && index >= start {
			currLength := i - start
			length = max(length, currLength)
			start = index + 1
		}
		indexes[ch] = i
	}
	length = max(length, len(s)-start)
	return length
}

func max(val1 int, val2 int) int {
	if val1 >= val2 {
		return val1
	}
	return val2
}

// func main() {
// 	s := "au"
// 	fmt.Println(lengthOfLongestSubstring(s))
// }
