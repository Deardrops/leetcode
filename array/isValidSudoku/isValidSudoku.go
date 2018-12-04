package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
	var dict map[byte]bool
	for i := 0; i < 9; i++ {
		dict = make(map[byte]bool)
		// 检查每行
		for j := 0; j < 9; j++ {
			num := board[i][j]
			if _, ok := dict[num]; ok && num != '.' {
				return false
			} else {
				dict[num] = true
			}
		}
		// 检查每列
		dict = make(map[byte]bool)
		for j := 0; j < 9; j++ {
			num := board[j][i]
			if _, ok := dict[num]; ok && num != '.' {
				return false
			} else {
				dict[num] = true
			}
		}
	}
	// 检查九宫格
	for offsetX := 0; offsetX < 9; offsetX = offsetX + 3 {
		for offsetY := 0; offsetY < 9; offsetY = offsetY + 3 {
			dict = make(map[byte]bool)
			for x := offsetX; x < offsetX+3; x++ {
				for y := offsetY; y < offsetY+3; y++ {
					num := board[x][y]
					if _, ok := dict[num]; ok && num != '.' {
						return false
					} else {
						dict[num] = true
					}
				}
			}
		}
	}

	return true
}

func main() {
	test := [][]byte{
		{'.', '.', '.', '.', '5', '.', '.', '1', '.'},
		{'.', '4', '.', '3', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '3', '.', '.', '1'},
		{'8', '.', '.', '.', '.', '.', '.', '2', '.'},
		{'.', '.', '2', '.', '7', '.', '.', '.', '.'},
		{'.', '1', '5', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '2', '.', '.', '.'},
		{'.', '2', '.', '9', '.', '.', '.', '.', '.'},
		{'.', '.', '4', '.', '.', '.', '.', '.', '.'},
	}
	res := isValidSudoku(test)
	fmt.Println(res)
}
