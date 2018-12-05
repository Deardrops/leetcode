package main

func strStr(haystack string, needle string) int {
	if len(haystack) == 0 {
		if len(needle) == 0 {
			return 0
		} else {
			return -1
		}
	}
outer:
	for i := 0; i < len(haystack); i++ {
		j := 0
		for ; j < len(needle); j++ {
			if i+j == len(haystack) {
				return -1
			}
			if haystack[i+j] != needle[j] {
				continue outer
			}
		}
		if j == len(needle) {
			return i
		}
	}
	return -1
}
