package main

import "fmt"

func generate(numRows int) [][]int {
	var res [][]int
	if numRows == 0 {
		return res
	}
	res = append(res, []int{1})
	if numRows == 1 {
		return res
	}
	res = append(res, []int{1, 1})
	if numRows == 2 {
		return res
	}
	for i := 2; i < numRows; i++ {
		var tmp []int
		tmp = append(tmp, 1)
		for j := 0; j+1 < len(res[i-1]); j++ {
			tmp = append(tmp, res[i-1][j]+res[i-1][j+1])
		}
		tmp = append(tmp, 1)
		res = append(res, tmp)
	}
	return res
}

func main() {
	i := generate(5)
	fmt.Println(i)
}
