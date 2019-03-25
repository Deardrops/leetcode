package main

import (
	"reflect"
	"strconv"
	"testing"
)

type input struct {
	haystack string
	needle   string
}

var tests = []struct {
	in  input
	out int
}{
	{input{"hello", "ll"}, 2},
	{input{"aaaaa", "bba"}, -1},
	{input{"", ""}, 0},
	{input{"", "j"}, -1},
	{input{"aaa", "aaaa"}, -1},
}

func Test(t *testing.T) {
	for i, tt := range tests {
		t.Run(strconv.Itoa(i+1), func(t *testing.T) {
			res := strStr(tt.in.haystack, tt.in.needle)
			if !reflect.DeepEqual(res, tt.out) {
				t.Errorf("with [%v] | got [%v] | want [%v]", tt.in, res, tt.out)
			}
		})
	}
}
