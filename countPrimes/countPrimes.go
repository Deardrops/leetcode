package main

/*
	Sieve of Eratosthenes
*/

func countPrimes(n int) int {
	if n < 2 {
		return 0
	}
	table := make([]bool, n+1)
	table[0], table[1] = true, true
	res := 0
	for i := 2; i < n; i++ {
		if table[i] {
			continue
		}
		res++
		for j := i * i; j < n; j += i {
			table[j] = true
		}
	}
	return res
}
