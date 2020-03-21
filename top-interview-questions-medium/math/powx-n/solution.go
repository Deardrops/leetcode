package main

import "fmt"

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		n = -n
		x = 1/x
	}
	return fastPow(x, n)
}

func fastPow(x float64, n int) float64 {
	if n == 1 {
		return x
	}
	tmp := fastPow(x, n/2)
	if n % 2 == 0 {
		return tmp * tmp
	} else {
		return tmp * tmp * x
	}
}
