package main

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	for i, j := 0, len(matrix[0])-1; i < len(matrix) && j >= 0; {
		switch {
		case matrix[i][j] == target:
			return true
		case matrix[i][j] > target:
			j--
		case matrix[i][j] < target:
			i++
		}
	}
	return false
}
