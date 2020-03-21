package main

func titleToNumber(s string) int {
	result := 0
	for _, char := range s {
        result = result * 26 + int(char - 'A' + 1)
	}
	return result
}