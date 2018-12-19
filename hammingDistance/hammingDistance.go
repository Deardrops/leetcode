package main

import "fmt"

func hammingDistance(x int, y int) int {
	var count int
	for b := x ^ y; b > 0; b >>= 1 {
		if b&1 == 1 {
			count++
		}
	}
	return count
}

func main() {
	res := hammingDistance(93, 73)
	fmt.Println(res)
}
