package main

import (
	"fmt"
)

func reverseBits(num uint32) uint32 {
	var res uint32
	for i := 31; i >= 0; i-- {
		res = res | ((num & 1) << uint(i))
		num = num >> 1
	}
	return res
}

func main() {
	var num uint32
	num = 43261596
	res := reverseBits(num)
	fmt.Println(res)
}
