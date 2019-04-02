package main

import "fmt"

/*
	解题思路：题意 = 每天只要比前一天的价格低，就卖出
*/
func maxProfit(prices []int) int {
	sum := 0
	for i := 1; i < len(prices); i++ {
		diff := prices[i] - prices[i-1]
		if diff > 0 {
			sum += diff
		}
	}
	return sum
}

func main() {
	prices := []int{7, 1, 5, 6, 3, 6}
	res := maxProfit(prices)
	fmt.Println(res)
}
