package main

import (
	"strconv"
)

func reverse(x int) int {
	str := strconv.Itoa(x)
	b := []byte(str)
	start := 0
	if str[0] == '-' {
		start = 1
	}
	for i, j := start, len(b)-1; i < j; i, j = i+1, j-1 {
		b[j], b[i] = b[i], b[j]
	}
	if start == 1 {
		b[0] = '-'
	}
	str = string(b)
	var num int
	var err error
	if num, err = strconv.Atoi(str); err != nil {
		return 0
	}
	if num > -2147483649 && num < 2147483648 {
		return num
	} else {
		return 0
	}
}
