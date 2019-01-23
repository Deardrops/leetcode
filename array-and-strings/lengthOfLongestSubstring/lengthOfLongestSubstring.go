package main

import "fmt"

// 方法一：用map在存储子串，key为字母，value为字母在子串中的序号（从0开始）
// 当扫描到的字符不在子串中时，在map中添加这个字符，并使最大长度加1
// 在子串中时，删除该字符，子串中所有字符的序号向前移动，再重新添加该字符
// 直到扫描完所有字符，返回结果
func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	subStrMap := make(map[int32]int)
	for _, char := range s {
		offset, ok := subStrMap[char]
		if !ok {
			subStrMap[char] = len(subStrMap)
			if len(subStrMap) > maxLen {
				maxLen++
			}
			continue
		}
		delete(subStrMap, char)
		for char := range subStrMap {
			tmp := subStrMap[char] - offset - 1
			if tmp < 0 {
				delete(subStrMap, char)
			} else {
				subStrMap[char] = tmp
			}
		}
		subStrMap[char] = len(subStrMap)
	}
	return maxLen
}

// 改进：value为字符串在子串中的序号的话，比较复杂，可以将value改为字母在原字符串中的位置
// source: https://leetcode.com/problems/longest-substring-without-repeating-characters/discuss/2158
func lengthOfLongestSubstring2(str string) int {
	m, max, left := make(map[rune]int), 0, 0
	for idx, c := range str {
		if _, okay := m[c]; okay == true && m[c] >= left {
			if idx-left > max {
				max = idx - left
			}
			left = m[c] + 1
		}
		m[c] = idx
	}
	if len(str)-left > max {
		max = len(str) - left
	}
	return max
}

func main() {
	str := "gaaqfeqlqky"
	res := lengthOfLongestSubstring(str)
	fmt.Println(res)
}
