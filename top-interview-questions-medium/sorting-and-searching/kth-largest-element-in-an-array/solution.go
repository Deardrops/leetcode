package main

import "fmt"

// 冒泡排序
func findKthLargest(nums []int, k int) int {
	for i := 0; i < len(nums); i++ {
		max := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > nums[max] {
				max = j
			}
		}
		if i+1 == k {
			return nums[max]
		}
		nums[i], nums[max] = nums[max], nums[i]
	}
	return 0
}

// TODO: 最小堆

func main() {
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	res := findKthLargest(nums, 4)
	fmt.Println(res)
}
