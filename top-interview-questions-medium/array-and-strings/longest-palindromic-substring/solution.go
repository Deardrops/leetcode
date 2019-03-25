package main

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	// even 代表以第i个字符为中心对称的子字符串（奇数个字符）
	// odd 代表以第i和第i+1个字符为中心对称的子字符串（偶数个字符）
	var even, odd string
	longest := s[0:1]
	for i := range s {
		left, right := i, i+1
		var stopCompareEven, stopCompareOdd bool
		for left >= 0 && right < len(s) {
			// 奇数个字符的比较
			if !stopCompareEven && s[left] == s[right] {
				even = s[left : right+1]
			} else {
				stopCompareEven = true
			}
			// 偶数个字符的比较
			if left-1 >= 0 && !stopCompareOdd && s[left-1] == s[right] {
				odd = s[left-1 : right+1]
			} else {
				stopCompareOdd = true
			}
			if stopCompareEven && stopCompareOdd {
				break
			}
			left--
			right++
		}
		// 选取 even 和 odd 中较长的一个
		if len(even) > len(longest) {
			longest = even
		}
		if len(odd) > len(longest) {
			longest = odd
		}
	}
	return longest
}
