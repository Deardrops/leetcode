package main

import (
	"reflect"
	"strconv"
	"testing"
)

var tests = []struct {
	in  []int
	out bool
}{
	{[]int{4, 5, 0, 1, 3}, true},
	{[]int{2, 5, 3, 4, 5}, true},
	{[]int{1, 2, 1, 2, 1, 2, 1, 2, 1, 2}, false},
}

func Test(t *testing.T) {
	for i, tt := range tests {
		t.Run(strconv.Itoa(i+1), func(t *testing.T) {
			res := increasingTriplet(tt.in)
			if !reflect.DeepEqual(res, tt.out) {
				t.Errorf("with [%v] | got [%v] | want [%v]", tt.in, res, tt.out)
			}
		})
	}
}
