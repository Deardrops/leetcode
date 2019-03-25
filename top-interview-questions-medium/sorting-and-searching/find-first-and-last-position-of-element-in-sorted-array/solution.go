package main

import "fmt"

func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}
	if len(nums) == 0 {
		return res
	}
	binarySearch(nums, 0, len(nums)-1, target, res)
	return res
}

func binarySearch(nums []int, low, high, target int, res []int) {
	if nums[low] == target {
		if low < res[0] || res[0] == -1 {
			res[0] = low
		}
	}
	if nums[high] == target && high > res[1] {
		res[1] = high
	}
	if low == high {
		return
	}
	mid := (low + high) / 2
	binarySearch(nums, low, mid, target, res)
	binarySearch(nums, mid+1, high, target, res)
}

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	res := searchRange(nums, 8)
	fmt.Println(res)
}
