package main

import "fmt"

func increasingTriplet(nums []int) bool {
	if len(nums) <= 2 {
		return false
	}
	var i, min, newMin, max int
	newMin = -1
	for i = 1; i < len(nums); i++ {
		if nums[i] > nums[min] {
			max = i
			break
		} else {
			min = i
		}
	}
	for i = i + 1; i < len(nums); i++ {
		if nums[i] > nums[max] {
			return true
		}
		if newMin < min {
			if nums[i] < nums[min] {
				newMin = i
				continue
			}
			if nums[i] > nums[min] && nums[i] < nums[max] {
				max = i
			}
		} else {
			if nums[i] < nums[newMin] {
				newMin = i
				continue
			}
			if nums[i] > nums[newMin] && nums[i] < nums[max] {
				min = newMin
				max = i
			}
		}
	}
	return false
}

func main() {
	nums := []int{1, 2, 1, 2, 1, 2, 1, 2}
	res := increasingTriplet(nums)
	fmt.Println(res)
}
