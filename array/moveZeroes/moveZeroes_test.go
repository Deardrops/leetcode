package main

import (
	"reflect"
	"strconv"
	"testing"
)

var tests = []struct {
	in  []int
	out []int
}{
	{[]int{0}, []int{0}},
	{[]int{1}, []int{1}},
	{[]int{1, 0}, []int{1, 0}},
	{[]int{0, 1}, []int{1, 0}},
	{[]int{2, 1}, []int{2, 1}},
	{[]int{1, 0, 1}, []int{1, 1, 0}},
	{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
}

func TestMoveZeros(t *testing.T) {
	baseTest(moveZeros, t)
}

func TestMoveZeros2(t *testing.T) {
	baseTest(moveZeros2, t)
}

func baseTest(f func([]int), t *testing.T) {
	for i, tt := range tests {
		t.Run(strconv.Itoa(i+1), func(t *testing.T) {
			f(tt.in)
			if !reflect.DeepEqual(tt.in, tt.out) {
				t.Errorf("got %v, want %v", tt.in, tt.out)
			}
		})
	}
}
