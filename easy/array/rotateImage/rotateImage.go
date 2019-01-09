package main

import "fmt"

func rotate(matrix [][]int) {
	n := len(matrix)
	for x := 0; x < n/2+1; x++ {
		for y := x; y < n-1-x; y++ {
			tmp := matrix[x][y]
			matrix[x][y] = matrix[n-y-1][x]
			matrix[n-y-1][x] = matrix[n-x-1][n-y-1]
			matrix[n-x-1][n-y-1] = matrix[y][n-x-1]
			matrix[y][n-x-1] = tmp
		}
	}
}

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	matrix2 := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	rotate(matrix)
	rotate(matrix2)
	fmt.Println(matrix)
	fmt.Println(matrix2)
}
