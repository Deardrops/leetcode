package main

import "fmt"

func findPeakElement(nums []int) int {
	return binarySearch(nums, 0, len(nums)-1)
}

func binarySearch(nums []int, low, high int) int {
	if low == high {
		return high
	}
	mid1 := (low + high) / 2
	mid2 := mid1 + 1
	if nums[mid1] > nums[mid2] {
		return binarySearch(nums, low, mid1)
	} else {
		return binarySearch(nums, mid2, high)
	}
}

func findPeakElement2(nums []int) int {
	var n int = len(nums)

	if n == 1 {
		return 0
	} else if n == 2 {
		if nums[0] > nums[1] {
			return 0
		} else {
			return 1
		}
	}

	var m int = n / 2
	if nums[m] > nums[m+1] && nums[m] > nums[m-1] {
		return m
	} else if nums[m-1] > nums[m] {
		return findPeakElement(nums[0:m])
	} else {
		return m + findPeakElement(nums[m+1:n]) + 1
	}

}

func main() {
	nums := []int{1, 2, 1, 3, 5, 6, 4}
	res := findPeakElement(nums)
	fmt.Println(res)
}
