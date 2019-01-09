package main

import "fmt"

func rotateArray(nums []int, k int) {
	n := len(nums)
	count := 0
	for start := 0; count < n; start++ {
		current := start
		prev := nums[start]
		for {
			next := (current + k) % n
			tmp := nums[next]
			nums[next] = prev
			prev = tmp
			current = next
			count++
			if start == current {
				break
			}
		}
	}
}

func main() {
	var nums = []int{1, 2}
	rotateArray(nums, 3)
	fmt.Printf("%v", nums)
}
