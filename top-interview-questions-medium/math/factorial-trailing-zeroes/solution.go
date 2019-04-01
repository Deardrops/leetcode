package main

/*
	解题思路：
*/
func trailingZeroes(n int) int {
	n = n / 5
	if n == 0 {
		return 0
	} else {
		return trailingZeroes(n) + n
	}
}
