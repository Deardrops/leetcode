package main

import "fmt"

func intersect(nums1 []int, nums2 []int) []int {
	var res []int
	dict := make(map[int]int)
	for _, i := range nums1 {
		dict[i]++
	}
	for _, j := range nums2 {
		if _, ok := dict[j]; ok {
			if dict[j] != 0 {
				res = append(res, j)
				dict[j]--
			}
		}
	}
	return res
}

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	res := intersect(nums1, nums2)
	fmt.Println(res)
}
