package main

import "fmt"

func majorityElement(nums []int) int {
	curr := -1
	count := 0
	for _, num := range nums {
		if count == 0 {
			curr = num
			count++
			continue
		}
		if num == curr {
			count++
		} else {
			count--
		}
	}
	return curr
}

func main() {
	input := []int{2,2,1,1,1,2,2}
	res := majorityElement(input)
	fmt.Println(res)
}