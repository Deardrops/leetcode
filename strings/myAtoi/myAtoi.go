package main

import "math"

func myAtoi(str string) int {
	strLen := len(str)
	if strLen == 0 {
		return 0
	}
	INT_MIN := -2147483648
	INT_MAX := 2147483647
	var numArray []int
	var c uint8
	var ok bool
	isNegative := false
	startIndex := 0
	for i, c := range str {
		if c != ' ' {
			startIndex = i
			break
		}
	}
	if c, ok = tryGetChar(startIndex, str); !ok {
		return 0
	}
	if c == '-' {
		isNegative = true
		startIndex++
	}
	if c == '+' {
		startIndex++
	}
	if c, ok = tryGetChar(startIndex, str); !ok {
		return 0
	}
	if !(48 <= c && c <= 57) {
		return 0
	}
	var i int
	for i = startIndex; i < len(str); i++ {
		if str[i] != '0' {
			startIndex = i
			break
		}
	}
	if i >= len(str) {
		return 0
	}
	for i := startIndex; i < len(str); i++ {
		c := str[i]
		if 48 <= c && c <= 57 {
			numArray = append(numArray, int(c-48))
		} else {
			break
		}
	}
	var res int
	length := len(numArray)
	if length > 10 {
		if isNegative {
			return INT_MIN
		} else {
			return INT_MAX
		}
	}

	for i := length - 1; i >= 0; i-- {
		multiplier := int(math.Pow10(length - i - 1))
		if isNegative {
			res -= numArray[i] * multiplier
			if res < INT_MIN {
				return INT_MIN
			}
		} else {
			res += numArray[i] * multiplier
			if res > INT_MAX {
				return INT_MAX
			}
		}
	}
	return res
}

func tryGetChar(i int, str string) (uint8, bool) {
	if i <= len(str)-1 {
		return str[i], true
	} else {
		return 0, false
	}
}

func myAtoi2(s string) int {
	characters := []byte(s)
	sign := 1
	num := 0
	inConversion := false
	for i := 0; i < len(characters); i++ {
		c := characters[i]
		if c == ' ' && inConversion == false {
			continue
		}

		if inConversion == false {
			if c == '+' {
				inConversion = true
			} else if c == '-' {
				inConversion = true
				sign = -1
			} else if v, ok := toDigit(c); ok {
				num = num*10 + v
				inConversion = true
			} else {
				return 0
			}
		} else {
			if v, ok := toDigit(c); ok {
				num = num*10 + v
				numWithSign := num * sign
				if numWithSign > math.MaxInt32 {
					return math.MaxInt32
				} else if numWithSign < math.MinInt32 {
					return math.MinInt32
				}
			} else {
				break
			}
		}
	}
	return num * sign
}

func toDigit(c byte) (int, bool) {
	if '0' <= c && c <= '9' {
		return int(c - '0'), true
	}
	return 0, false
}
