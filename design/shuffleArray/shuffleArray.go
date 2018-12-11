package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	解题方法：洗牌算法（Fisher Yates shuffle）
*/

type Solution struct {
	nums   []int
	random *rand.Rand
}

/** Returns a random shuffling of the array. */

func Constructor(nums []int) Solution {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	return Solution{nums, random}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.nums
}

func (this *Solution) Shuffle() []int {
	length := len(this.nums)
	nums := make([]int, length)
	copy(nums, this.nums)
	for i := length - 1; i >= 0; i-- {
		j := this.random.Intn(i + 1)
		nums[i], nums[j] = nums[j], nums[i]
	}
	return nums
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
	obj := Constructor(nums)
	for i := 0; i < 100; i++ {
		res := obj.Shuffle()
		fmt.Println(res)
	}
}
