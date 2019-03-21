package main

func uniquePaths(m int, n int) int {
	result := make([][]int, m)
	// 初始化一个二维数组，大小和题目中给的一样，每个元素都是0
	for i := 0; i < m; i++ {
		result[i] = make([]int, n)
	}

	// 将二维数组第一行全部刷为0
	for i := 0; i < m; i++ {
		result[i][0] = 1
	}

	// 将二维数组第一列全部刷为0
	for j := 1; j < n; j++ {
		result[0][j] = 1
	}

	// 每个元素的值代表走到这一格的可能情况数量
	// 可能情况数量 = 走到它左边的可能情况 + 走到它上面的可能情况数量
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			result[i][j] = result[i-1][j] + result[i][j-1]
		}
	}

	// 返回最后一个元素的可能情况数量
	return result[m-1][n-1]
}
