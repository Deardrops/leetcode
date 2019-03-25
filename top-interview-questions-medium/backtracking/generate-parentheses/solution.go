package main

import "fmt"

// Copied from https://leetcode.com/problems/generate-parentheses/discuss/172719
// n对括号，有n个左括号和n个右括号
// 从左向右挨个添加括号，先添加左括号，然后是右括号，这里使得每次迭代字符串一分为二
// 如果两种括号已经消耗完的话，添加到结果中
// 如果左括号未消耗完的话，添加一个左括号，并拷贝一份新的数组，继续递归
// 若果右括号未消耗完的话，添加一个右括号，并拷贝一份新的数组，继续递归
// 总结一句，我没有从括号的数量这个角度出发考虑，也想不出这个一分为二的写法来

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}

	leftRemain, rightRemain := n-1, n
	cur := []byte{'('}
	var res []string
	doGenerateParenthesis(leftRemain, rightRemain, cur, &res)
	return res
}

func doGenerateParenthesis(leftRemain, rightRemain int, cur []byte, res *[]string) {
	if leftRemain == 0 && rightRemain == 0 {
		*res = append(*res, string(cur))
		return
	}

	if leftRemain > 0 {
		doGenerateParenthesis(leftRemain-1, rightRemain, append(cur, '('), res)
	}
	if rightRemain > leftRemain {
		doGenerateParenthesis(leftRemain, rightRemain-1, append(cur, ')'), res)
	}
}

func main() {
	res := generateParenthesis(3)
	fmt.Println(res)
}
