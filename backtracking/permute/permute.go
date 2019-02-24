package main

import "fmt"

func permute(nums []int) [][]int {
	var ret [][]int
	recursion(nums, []int{}, &ret)
	return ret
}

func recursion(nums []int, curr []int, ret *[][]int) {
	if len(nums) == 1 {
		curr = append(curr, nums[0])
		*ret = append(*ret, curr)
		return
	}

	for i, num := range nums {
		tmp := append([]int{}, nums[:i]...)
		tmp = append(tmp, nums[i+1:]...)
		recursion(tmp, append(curr, num), ret)
	}
}

func main() {
	res := permute([]int{1, 2, 3})
	fmt.Println(res)
}
