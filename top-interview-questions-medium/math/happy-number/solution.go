package main

import "fmt"

func isHappy(n int) bool {
	for i := 0; i < 10; i++ {
		sum := 0
		for _, num := range IntToSlice(n) {
			sum += num * num
		}
		if sum == 1 {
			return true
		}
		n = sum
	}
	return false
}

func IntToSlice(n int) []int {
	var nums []int
	for {
		quotient := n / 10
		remainder := n % 10
		nums = append([]int{remainder}, nums...)
		if quotient == 0 {
			break
		}
		n = quotient
	}
	return nums
}

func main() {
	fmt.Println(isHappy(19))
}
