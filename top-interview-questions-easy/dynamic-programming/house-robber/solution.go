package main

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var tmp int
	prev1 := 0
	prev2 := 0
	for _, value := range nums {
		tmp = prev1
		if prev2+value > prev1 {
			prev1 = prev2 + value
		}
		prev2 = tmp
	}
	return prev1
}
