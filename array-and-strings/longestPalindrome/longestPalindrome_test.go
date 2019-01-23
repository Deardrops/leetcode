package main

import (
	"reflect"
	"strconv"
	"testing"
)

var tests = []struct {
	in  string
	out string
}{
	{"babad", "bab"},
	{"cbbd", "bb"},
	{"bbbb", "bbbb"},
	{"bbbbb", "bbbbb"},
}

func Test(t *testing.T) {
	for i, tt := range tests {
		t.Run(strconv.Itoa(i+1), func(t *testing.T) {
			res := longestPalindrome(tt.in)
			if !reflect.DeepEqual(res, tt.out) {
				t.Errorf("with [%v] | got [%v] | want [%v]", tt.in, res, tt.out)
			}
		})
	}
}
