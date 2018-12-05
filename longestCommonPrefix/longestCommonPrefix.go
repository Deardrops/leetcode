package main

func longestCommonPrefix(strs []string) string {
	var res []byte
	if len(strs) == 0 {
		return ""
	}
outer:
	for i := 0; ; i++ {
		if i >= len(strs[0]) {
			break
		}
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) {
				break outer
			}
			if strs[j][i] != strs[j-1][i] {
				break outer
			}
		}
		res = append(res, strs[0][i])
	}
	return string(res)
}
