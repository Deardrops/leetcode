package main

import "fmt"

func singleNumber(nums []int) int {
	res := 0

	for _, num := range nums {
		res ^= num
	}

	return res
}

func main() {
	testSet := []int{4, 1, 2, 1, 2}
	res := singleNumber(testSet)
	fmt.Println(res)
}
