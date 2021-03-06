package main

import "fmt"

func exist(board [][]byte, word string) bool {
	height := len(board)
	if height == 0 {
		return false
	}
	width := len(board[0])
	if width == 0 {
		return false
	}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			var visited [][]bool
			for i := 0; i < height; i++ {
				tmp := make([]bool, width)
				visited = append(visited, tmp)
			}
			if search(width, height, x, y, []byte(word), board, visited) {
				return true
			}
		}
	}
	return false
}

func search(width, height, x, y int, str []byte, board [][]byte, visited [][]bool) bool {
	if len(str) == 0 {
		return true
	}
	if x < 0 || x >= width || y < 0 || y >= height {
		return false
	}
	if str[0] != board[y][x] {
		return false
	}
	if visited[y][x] {
		return false
	}

	visited[y][x] = true
	str = str[1:]
	succ := search(width, height, x, y-1, str, board, visited) || search(width, height, x, y+1, str, board, visited) || search(width, height, x-1, y, str, board, visited) || search(width, height, x+1, y, str, board, visited)
	// Q：为什么这里要将已修改的二维数组还原回去？按理说数组在上一句函数调用时会拷贝一份，原数组并不会被修改才对。
	// A：《Go 语言学习笔记》中提到以下几点：
	// []int 这个类型叫切片，不叫数组。[3]int 这种才叫做数组。
	// 切片是一种特殊的结构体，其内部通过指针引用底层数组，可以把切片看作数组指针的一种封装。
	// Golang 语言中，函数的参数一定是值拷贝传递（pass-by-value）。
	// 新建的切片对象依旧指向原底层数组，也就是说修改度所有关联的切片可见。
	// 两句话加起来是说，切片作为参数时，在函数内部被修改时，会影响到原来的切片。

	if !succ {
		visited[y][x] = false
	} else {
		return true
	}
	return false
}

// 大佬的解法，不需要单独开二维数组保存访问过的位置
// 直接用一个函数内的局部变量保存当前访问位置的字符，然后将这个位置的元素设置为一个标志字符（比如“0”）
// 当向下搜索失败，回溯到这个函数的时候，再讲这个局部变量保存的字符还回原来的位置
// 这样做节省了不知道多少代码指数和时间
func exist2(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == word[0] {
				if backtrackExist(board, i, j, word[1:]) {
					return true
				}
			}
		}
	}
	return false
}

func backtrackExist(board [][]byte, i, j int, w string) bool {
	if w == "" {
		return true
	}
	// set board[i][j] empty
	c := board[i][j]
	board[i][j] = '1'
	// check (i-1, j)
	if i > 0 && board[i-1][j] == w[0] {
		if backtrackExist(board, i-1, j, w[1:]) {
			return true
		}
	}
	// check (i+1, j)
	if i < len(board)-1 && board[i+1][j] == w[0] {
		if backtrackExist(board, i+1, j, w[1:]) {
			return true
		}
	}
	// check (i, j-1)
	if j > 0 && board[i][j-1] == w[0] {
		if backtrackExist(board, i, j-1, w[1:]) {
			return true
		}
	}
	if j < len(board[0])-1 && board[i][j+1] == w[0] {
		if backtrackExist(board, i, j+1, w[1:]) {
			return true
		}
	}
	board[i][j] = c
	return false
}

func main() {
	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'E', 'S'}, {'A', 'D', 'E', 'E'}}
	word := "ABCESEEEFS"
	fmt.Println(exist(board, word))
}
