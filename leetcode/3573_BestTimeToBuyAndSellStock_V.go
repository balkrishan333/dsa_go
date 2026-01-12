package main

func maximumProfit(prices []int, k int) int64 {
	days := len(prices)
	dp := make([][][]int64, days)

	for i := range dp {
		dp[i] = make([][]int64, k+1)
		for j := range dp[i] {
			dp[i][j] = make([]int64, 3)
		}
	}

	//initialize base cases
	for tx := 1; tx <= k; tx++ {
		dp[0][tx][0] = 0
		dp[0][tx][1] = -int64(prices[0])
		dp[0][tx][2] = int64(prices[0])
	}

	for day := 1; day < days; day++ {
		for tx := 1; tx <= k; tx++ {
			price := int64(prices[day])
			dp[day][tx][0] = max(dp[day-1][tx][0], dp[day-1][tx][1]+price, dp[day-1][tx][2]-price)
			dp[day][tx][1] = max(dp[day-1][tx][1], dp[day-1][tx-1][0]-price)
			dp[day][tx][2] = max(dp[day-1][tx][2], dp[day-1][tx-1][0]+price)
		}
	}
	return int64(dp[days-1][k][0])
}

// func main() {
// 	println(maximumProfit([]int{1, 7, 9, 8, 2}, 2))
// }
