package main

// 标准 DP 题目
func coinChange(coins []int, amount int) int {
	// 初始化一个空数组，其中每个元素为达到当前金额所需要的最少硬币数量
	dp := make([]int, amount+1)
	// 这里有个技巧化的处理，最小的硬币为1，我们将每个元素初始值都设为目标值+1，而实际的硬币数量绝不会超过这个值
	for v := 1; v <= amount; v++ {
		dp[v] = amount + 1
	}
	for _, coin := range coins {
		// 对指定的硬币，尝试添加进来，如果添加后的硬币总数量少于当前值，则替换当前值
		for v := coin; v <= amount; v++ {
			if dp[v] > dp[v-coin]+1 {
				dp[v] = dp[v-coin] + 1
			}
		}
	}
	// 如果目标值的位置，没有任何情况满足，则返回-1
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}
