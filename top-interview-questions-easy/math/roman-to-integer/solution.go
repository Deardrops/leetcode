package main

func romanToInt(s string) int {
	res := 0
	var sumC, sumX, sumI int
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'M':
			res += 1000
			if sumC != 0 {
				res -= sumC
				sumC = 0
			}
		case 'D':
			res += 500
			if sumC != 0 {
				res -= sumC
				sumC = 0
			}
		case 'C':
			sumC += 100
			if sumX != 0 {
				res -= sumX
				sumX = 0
			}
		case 'L':
			res += 50
			if sumX != 0 {
				res -= sumX
				sumX = 0
			}
		case 'X':
			sumX += 10
			if sumI != 0 {
				res -= sumI
				sumI = 0
			}
		case 'V':
			res += 5
			if sumI != 0 {
				res -= sumI
				sumI = 0
			}
		case 'I':
			sumI += 1
		}
	}
	res = res + sumC + sumX + sumI
	return res
}
