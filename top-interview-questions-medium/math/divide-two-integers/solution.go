package main

// 1. 将除数和被除数都转换为帧数
// 2. 用位运算加速（替代乘法）
// 3. 最后处理数值越界问题
func divide(dividend int, divisor int) int {
	var symbol1 bool
	var symbol2 bool
	if dividend < 0 {
		symbol1 = true
		dividend = -dividend
	}
	if divisor < 0 {
		symbol2 = true
		divisor = -divisor
	}
	res := 0
	for dividend >= divisor {
		tmp, i := divisor, 1
		for dividend >= tmp {
			dividend -= tmp
			res += i
			i <<= 1 // i = i + i
			tmp <<= 1 // tmp = tmp + tmp
		}
	}
	if (!symbol1 && symbol2) || (symbol1 && !symbol2) {
		res = -res
	}
	max := 1 << 31 - 1
	min := - (1 << 31)

	if res > max {
		return max
	}
	if res < min {
		return min
	}
	return res
}