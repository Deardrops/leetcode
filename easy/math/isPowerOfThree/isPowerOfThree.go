package main

func isPowerOfThree(n int) bool {
	i := float64(n)
	for i >= 3 {
		i = i / 3.0
	}
	if i == 1 {
		return true
	} else {
		return false
	}
}
