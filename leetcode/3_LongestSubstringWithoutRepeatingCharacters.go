package main

func lengthOfLongestSubstring(s string) int {
	symbols := [255]int{}
	for i := range symbols {
		symbols[i] = -1
	}

	start, maxlen := 0, 0
	for pos, char := range s {
		//fmt.Printf("char %c pos %d\n", char, pos)
		if symbols[char] != -1 {
			if symbols[char] >= start {
				if maxlen < pos-start {
					maxlen = pos - start
				}
				start = symbols[char] + 1
			}
		}
		symbols[char] = pos
	}

	if maxlen < len(s)-start {
		maxlen = len(s) - start
	}
	return maxlen
}

// func main() {
// 	fmt.Print("Max len : ", lengthOfLongestSubstring("dvdf"), "\n")
// }
