package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	meadianIndex := (len(nums1) + len(nums2)) / 2
	mergeArr := make([]int, meadianIndex+1)
	var i, j int = 0, 0

	for currIndex := 0; currIndex <= meadianIndex; currIndex++ {
		if i < len(nums1) && j < len(nums2) {
			if nums1[i] <= nums2[j] {
				mergeArr[currIndex] = nums1[i]
				i++
			} else {
				mergeArr[currIndex] = nums2[j]
				j++
			}
		} else if i < len(nums1) {
			mergeArr[currIndex] = nums1[i]
			i++
		} else if j < len(nums2) {
			mergeArr[currIndex] = nums2[j]
			j++
		}
	}

	if (len(nums1)+len(nums2))%2 != 0 {
		return float64(mergeArr[meadianIndex])
	}

	return float64(mergeArr[meadianIndex]+mergeArr[meadianIndex-1]) / 2
}

// func main() {
// 	nums1 := []int{}
// 	nums2 := [1]int{3}

// 	fmt.Println(findMedianSortedArrays(nums1[:], nums2[:]))
// }
