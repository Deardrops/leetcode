package main

import "fmt"

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	counts := make([]int, len(nums))
	maxs := make([]int, len(nums))
	max := 1

	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if maxs[j] < nums[i] && counts[j]+1 > counts[i] {
				maxs[i] = nums[i]
				counts[i] = counts[j] + 1
				if counts[i] > max {
					max = counts[i]
				}
			}
		}
		if counts[i] == 0 {
			counts[i] = 1
			maxs[i] = nums[i]
		}
	}
	return max
}

func main() {
	//nums := []int{3,1,2}
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	res := lengthOfLIS(nums)
	fmt.Println(res)
}
