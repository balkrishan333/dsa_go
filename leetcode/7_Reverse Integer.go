package main

func reverse(x int) int {
	var result int
	var max = 1<<31 - 1
	var min = -1 << 31

	for x != 0 {
		result = result*10 + x%10
		x /= 10
	}
	if result > max || result < min {
		return 0
	}
	return result
}

func main() {
	println(reverse(-123))
}
