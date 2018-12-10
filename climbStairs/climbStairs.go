package main

/*
	解题思路：
	要爬到第i级楼梯，可以拆解成两个部分，
	①先爬到第i-1级楼梯，再爬1级楼梯，
	②先爬到第i-2级楼梯，再（一口气）爬2级楼梯
	注意，①和②两部分中没有重复方法，这里以用集合论的思想来理解，即这两个集合没有交集，
	因此，爬到第i级的方法总数=①的方法总数+②的方法总数，
	显然，①的方法总数就是爬到第i-1级楼梯的所有方法的总数，②同理，
	因此，建立递推式f(i) = f(i-1) + f(i-2)，可以求得最终结果。
*/

func climbStairs(n int) int {
	solution := make([]int, n+1)
	if n <= 3 {
		return n
	}
	solution[0] = 0
	solution[1] = 1
	solution[2] = 2
	for i := 3; i <= n; i++ {
		solution[i] = solution[i-1] + solution[i-2]
	}
	return solution[n]
}
