package main

import "fmt"

func firstBadVersion(n int) int {
	var mid, min, max int
	min = 0
	max = n
	for max != min {
		mid = (min + max) / 2
		if isBadVersion(mid) {
			max = mid
		} else {
			if mid == max-1 {
				return max
			}
			min = mid
		}
	}
	return 0
}

func main() {
	res := firstBadVersion(10)
	fmt.Println(res)
}

func isBadVersion(n int) bool {
	if n >= 6 {
		return true
	} else {
		return false
	}
}
