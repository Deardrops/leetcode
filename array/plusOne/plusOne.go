package main

import "fmt"

func plusOne(digits []int) []int {
	res := make([]int, len(digits))
	copy(res, digits)
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {
			if i == 0 {
				res[i] = 1
				res = append(res, 0)
			} else {
				res[i] = 0
			}
		} else {
			res[i]++
			break
		}
	}
	return res
}

func main() {
	digits := []int{1, 2, 3, 4}
	res := plusOne(digits)
	fmt.Println(res)
}
