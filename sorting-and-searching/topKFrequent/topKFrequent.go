package main

import "fmt"

func topKFrequent(nums []int, k int) []int {
	ret := make(map[int]int)
	for _, num := range nums {
		if _, ok := ret[num]; ok {
			ret[num]++
		} else {
			ret[num] = 1
		}
	}
	var res []int
outer:
	for key, value := range ret {
		for i, num := range res {
			if value > ret[num] {
				res = append(res[:i], append([]int{key}, res[i:]...)...)
				continue outer
			}
		}
		res = append(res, key)
	}
	return res[:k]
}

func main() {
	nums := []int{-1, 1, 4, -4, 3, 5, 4, -2, 3, -1}
	res := topKFrequent(nums, 3)
	fmt.Println(res)
}
