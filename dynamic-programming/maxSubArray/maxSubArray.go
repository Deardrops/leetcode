package main

func maxSubArray(nums []int) int {
	maxCurr := nums[0]
	maxSoFar := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i]+maxCurr {
			maxCurr = nums[i]
		} else {
			maxCurr += nums[i]
		}
		if maxCurr > maxSoFar {
			maxSoFar = maxCurr
		}
	}
	return maxSoFar
}
