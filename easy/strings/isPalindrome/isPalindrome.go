package main

import "fmt"

func convToLower(c uint8) (uint8, bool) {
	if c >= 48 && c <= 57 {
		return c, true
	}
	if c >= 65 && c <= 90 {
		return c + 32, true
	}
	if c >= 97 && c <= 122 {
		return c, true
	}
	return 0, false
}

func isPalindrome(s string) bool {
	var a, b uint8
	var ok bool
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		for ; i < j; i++ {
			if a, ok = convToLower(s[i]); ok {
				break
			}
		}
		if i == j {
			break
		}
		for ; i < j; j-- {
			if b, ok = convToLower(s[j]); ok {
				break
			}
		}
		if i == j {
			break
		}
		if b != a {
			return false
		}
	}
	return true
}

func main() {
	str := "a."
	res := isPalindrome(str)
	fmt.Println(res)
}
