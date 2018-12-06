package main

import (
	"reflect"
	"strconv"
	"testing"
)

type input struct {
	s string
	t string
}

var tests = []struct {
	in  input
	out bool
}{
	{input{"rat", "taa"}, false},
	{input{"anagram", "nagaram"}, true},
}

func Test(t *testing.T) {
	for i, tt := range tests {
		t.Run(strconv.Itoa(i+1), func(t *testing.T) {
			res := isAnagram(tt.in.s, tt.in.t)
			if !reflect.DeepEqual(res, tt.out) {
				t.Errorf("with [%v] | got [%v] | want [%v]", tt.in, res, tt.out)
			}
		})
	}
}
