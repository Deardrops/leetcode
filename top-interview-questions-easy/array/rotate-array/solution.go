package main

import "fmt"

// 方法一：用一个新数组存储所有元素旋转后的位置

// 方法二：迭代进行每个元素的移动
func rotate(nums []int, k int) {
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

// 方法三：三反转
func rotate3(nums []int, k int) {
	lenn := len(nums)
	if lenn == 0 {
		return
	}
	k = k % lenn
	var zz []int
	for i := 0; i < k; i++ {
		zz = append(zz, nums[i+lenn-k])
	}
	for i := 0; i < lenn-k; i++ {
		nums[lenn-1-i] = nums[lenn-1-i-k]
	}
	for i := 0; i < k; i++ {
		nums[i] = zz[i]
	}
}

func main() {
	var nums = []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	fmt.Printf("%v", nums)
}
