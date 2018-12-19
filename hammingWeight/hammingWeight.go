package main

/*
	解题思路：
	把num看作二进制数，用二进制运算符来操作它
	num & 1 会将两者进行且运算，返回的结果是1或者0
	即判断num的最后一位是否为1，将结果加到计数器上
	然后将num右移一位，继续下一次迭代
	最后计数器的变量值为所有1的个数
*/
func hammingWeight(num uint32) int {
	res := 0
	for num > 0 {
		res += int(num & 1)
		num = num >> 1
	}
	return res
}
