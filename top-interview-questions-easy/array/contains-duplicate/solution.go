package main

import "fmt"

func containsDuplicate(nums []int) bool {
	var arrayMap map[int]bool
	arrayMap = make(map[int]bool)
	for _, num := range nums {
		if _, ok := arrayMap[num]; ok {
			return true
		} else {
			arrayMap[num] = true
		}
	}
	return false
}

func main() {
	testSet := []int{1, 2, 3, 1}
	res := containsDuplicate(testSet)
	fmt.Println(res)
}
