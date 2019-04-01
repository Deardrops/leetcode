package main

import (
	"fmt"
	"math/rand"
)

type RandomizedSet struct {
	nums []int
	set  map[int]int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{
		nums: make([]int, 0),
		set:  make(map[int]int, 0),
	}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.set[val]; ok {
		return false
	} else {
		this.nums = append(this.nums, val)
		this.set[val] = len(this.nums) - 1
		return true
	}
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	if idx, ok := this.set[val]; ok {
		this.nums = append(this.nums[:idx], this.nums[idx+1:]...)
		delete(this.set, val)
		for num, i := range this.set {
			if i > idx {
				this.set[num] = i - 1
			}
		}
		return true
	} else {
		return false
	}
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Int()%len(this.nums)]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */

func main() {
	rs := Constructor()
	fmt.Println(rs.Insert(0))
	fmt.Println(rs.Insert(1))
	fmt.Println(rs.Remove(0))
	fmt.Println(rs.Insert(2))
	fmt.Println(rs.Remove(1))
	fmt.Println(rs.GetRandom())
}
