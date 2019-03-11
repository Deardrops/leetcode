package main

import "testing"

var nums = []int{1, 2, 1, 3, 5, 6, 4}

func BenchmarkS1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findPeakElement(nums)
	}
}

func BenchmarkS2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findPeakElement2(nums)
	}
}
