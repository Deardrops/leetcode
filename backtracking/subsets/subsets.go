package main

import "fmt"

func subsets(nums []int) [][]int {
	ret := [][]int{{}}
	recursion(&nums, 0, &ret)
	return ret
}

func recursion(nums *[]int, idx int, ret *[][]int) {
	length := len(*ret)
	for i := 0; i < length; i++ {
		tmp := append([]int{}, (*ret)[i]...)
		tmp = append(tmp, (*nums)[idx])
		*ret = append(*ret, tmp)
	}
	idx++
	if idx == len(*nums) {
		return
	}
	recursion(nums, idx, ret)
}

func main() {
	res := subsets([]int{9, 0, 3, 5, 7})
	fmt.Println(res)
}
