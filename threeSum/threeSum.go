package main

import (
	"fmt"
	"sort"
)

/*
	解题思路：
	先对数组排序，固定最左边的数字A，A右边的数组求解目标值为-A的twoSum的问题
	逐个移动右边的两个元素，如果包含A在内的三个数字之和为0，添加到结果中
	为了提高效率，需要在移动时前后比较，跳过相同的数字（具体移动方法见解法2）
*/
// 个人解法，效率很差，原因是没有考虑到：
// 如果三数之和已经大于0了，继续向右移动这些数求和比较也没有意义，不会产生正确的结果来
// 因此应该将这种情况跳过，进行下一阶段的循环
func threeSum(nums []int) [][]int {
	length := len(nums)
	var res [][]int
	if length < 3 {
		return res
	}
	sort.Ints(nums)
	for i := 0; i < length-2; i++ {
		for j := i + 1; j < length-1; j++ {
			for k := j + 1; k < length; k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					res = append(res, []int{nums[i], nums[j], nums[k]})
				}
				for k+1 < length && nums[k+1] == nums[k] {
					k++
				}

			}
			for j+1 < length-1 && nums[j+1] == nums[j] {
				j++
			}
		}
		for i+1 < length-2 && nums[i+1] == nums[i] {
			i++
		}
	}
	return res
}

// Leetcode 上的解法，beat 89%
func threeSum1(nums []int) [][]int {
	length := len(nums)
	solution := make([][]int, 0)
	if length < 3 {
		return solution
	}
	sort.Ints(nums) // 先排序
	// 正在移动中的三个的数字分别为 ABC（从左到右）
	for A := range nums {
		if A > 0 && nums[A] == nums[A-1] {
			continue // 移动A时，直接跳过相同的数值，移动到下一个不相同的数字处
		}
		B := A + 1      // 取数字A右边的一个数字
		C := length - 1 // 取数组的最后一个数字
		for B < C {     // 在两者相遇之前
			sum := nums[A] + nums[B] + nums[C]
			if sum == 0 { // 如果三数之和满足条件
				array := []int{nums[A], nums[B], nums[C]}
				solution = append(solution, array) // 添加到结果数组中
				B++                                // B向后移动一位
				C--                                // C向前移动一位
				for B < C && nums[B] == nums[B-1] {
					B++ // B移动到下一个不同的数字处
				}
				for B < C && nums[C] == nums[C+1] {
					C-- // C移动到上一个不同的数字处
				}
			} else if sum > 0 {
				C-- //如果所求之和大了，则用更小的C尝试
			} else if sum < 0 {
				B++ // 如果所求之和小了，则用更大的B尝试
			}
		}
	}
	return solution
}

func main() {
	nums := []int{0, 0, 0}
	res := threeSum(nums)
	fmt.Println(res)
}
