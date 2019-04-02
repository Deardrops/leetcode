package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	i := 1
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[j-1] {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}

func main() {
	nums := []int{1, 1, 2}
	i := removeDuplicates(nums)
	fmt.Println(nums[:i])
}
