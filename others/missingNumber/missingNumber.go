package main

import "sort"

/*
	三种解法：
	1. 差异或 xor：用一个0到n完整的数组和原数组两个数组中的所有元素进行差异或运算，最后剩下的数字是缺失的那个数字
	2. 二分搜索：先对数组排序，再进行二分搜索，找到缺失的数字
	3. 求和法：已知数组的长度为 n，完整数组长度为n+1，完整数组的所有元素求和sum=n(n-1)/2，用该和减去数组所有元素的累加值，即为缺失的数字
*/

// 差异或，时间复杂度O(n)，空间复杂度O(1)
func missingNumber1(nums []int) int {
	nlen := len(nums)
	if nlen == 0 {
		return 0
	}

	res := nums[0]
	for i := 0; i < nlen+1; i++ {
		if i == 0 || i == nlen {
			res = res ^ i
			continue
		}
		res = res ^ i ^ nums[i]
	}
	return res
}

// 二分搜索，时间复杂度O(nlog(n))，空间复杂度O(1)
func missingNumber2(nums []int) int {
	sort.Ints(nums)
	left, right := 0, len(nums)
	for left < right {
		mid := (left + right) / 2
		if nums[mid] > mid {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}

// 求和法，时间复杂度O(n)，空间复杂度O(1)
func missingNumber3(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	n := len(nums) + 1
	return n*(n-1)/2 - sum
}
