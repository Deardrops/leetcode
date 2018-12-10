package main

/*
	解题思路：
	每次迭代过程中，计算相邻两个数值的差值，并累加到最大差值上。
	如果累加后的最大差值小于0，将其重置为0。
	然后比较这个最大差值和历史最大差值，取相对较大的一个作为新的历史最大差值。
	迭代结束后，最终的历史最大差值即为求解结果。
*/

func maxProfit(prices []int) int {
	var maxCurr, maxDiff int
	for i := 1; i < len(prices); i++ {
		maxCurr += prices[i] - prices[i-1]
		if maxCurr < 0 {
			maxCurr = 0
		}
		if maxCurr > maxDiff {
			maxDiff = maxCurr
		}
	}
	return maxDiff
}
