package main

import "fmt"

var letterMap = map[uint8]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func letterCombinations(digits string) []string {
	var ret []string
	for _, char := range digits {
		var tmp []string
		if len(ret) == 0 {
			for _, letter := range letterMap[uint8(char)] {
				tmp = append(tmp, string(letter))
			}
		} else {
			for _, str := range ret {
				for _, letter := range letterMap[uint8(char)] {
					tmp = append(tmp, str+string(letter))
				}
			}
		}
		ret = tmp
	}
	return ret
}

func main() {
	res := letterCombinations("22")
	fmt.Println(res)
}
