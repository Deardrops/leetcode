package main

import "fmt"

/*
 	解题方法：
	columns 是列数， rows 是行数
	先行后列，左上角到右下角扫描每个元素，如果遇到为0的元素，将该行的第一个元素置0，该列的第一个元素置0
	为了避免第一列的元素被覆盖，单独用一个变量 col0 标识，表示第一列中有0的元素
	最后先行列后，从右下角到左上角扫描每个元素，如果该行的第一个元素和该列的第一个元素为0，将这个元素置为0
	col0 表示为 true 时，将第一列每个元素置为0
*/

func setZeroes(matrix [][]int) {
	rows := len(matrix)
	if rows == 0 {
		return
	}
	cols := len(matrix[0])
	col0 := false
	for i := 0; i < rows; i++ {
		if matrix[i][0] == 0 {
			col0 = true
		}
		for j := 1; j < cols; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := rows - 1; i >= 0; i-- {
		for j := cols - 1; j >= 1; j-- {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
		if col0 {
			matrix[i][0] = 0
		}
	}
}

func main() {
	matrix := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}
	setZeroes(matrix)
	fmt.Println(matrix)
}
