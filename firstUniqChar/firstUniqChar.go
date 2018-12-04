package main

import "fmt"

func firstUniqChar(s string) int {
	dict := make(map[int32]int)
	for _, c := range s {
		if _, ok := dict[c]; ok {
			dict[c]++
		} else {
			dict[c] = 1
		}
	}
	for i, c := range s {
		if dict[c] == 1 {
			return i
		}
	}
	return -1
}

func firstUniqChar1(s string) int {
	flags := make([]int, 26)
	for i := range flags {
		flags[i] = -1
	}
	slen := len(s)

	for i, ch := range s {
		idx := byte(ch - 'a')
		if flags[idx] == -1 {
			flags[idx] = i
		} else {
			flags[idx] = slen
		}
	}

	min := slen
	for i := range flags {
		if flags[i] >= 0 && flags[i] < len(s) && flags[i] < min {
			min = flags[i]
		}
	}

	if min == slen {
		return -1
	}
	return min
}

func main() {
	s := "loveleetcode"
	res := firstUniqChar(s)
	fmt.Println(res)
}
