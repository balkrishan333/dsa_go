package main

func twoSum(nums []int, target int) []int {

	indexMap := make(map[int]int)
	var answer [2]int

	for i := 0; i < len(nums); i++ {
		val, ok := indexMap[target-nums[i]]

		if ok {
			answer[0] = i
			answer[1] = val
		} else {
			indexMap[nums[i]] = i
		}
	}
	return answer[:]
}

// func main() {
// 	nums := [4]int{2, 7, 11, 15}
// 	target := 9
// 	answer := twoSum(nums[:], target)

// 	fmt.Println(answer)
// }
