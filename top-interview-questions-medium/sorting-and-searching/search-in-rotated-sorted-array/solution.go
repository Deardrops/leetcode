package main

import "fmt"

func search(nums []int, target int) int {
	return binarySearch(nums, 0, target)
}

func binarySearch(nums []int, start, target int) int {
	if len(nums) == 0 {
		return -1
	}
	low := 0
	high := len(nums) - 1
	mid := (low + high) / 2
	if nums[mid] == target {
		return start + mid
	}
	res := binarySearch(nums[:mid], start, target)
	if res != -1 {
		return res
	}
	res = binarySearch(nums[mid+1:], start+mid+1, target)
	return res
}

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	res := search(nums, 2)
	fmt.Println(res)
}
