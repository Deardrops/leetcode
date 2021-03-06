package main

import (
	"reflect"
	"strconv"
	"testing"
)

var tests = []struct {
	in  string
	out int
}{
	{"   -42", -42},
	{"4193 with words", 4193},
	{"words and 987", 0},
	{"-91283472332", -2147483648},
	{"3.14159", 3},
	{"+1", 1},
	{"+-2", 0},
	{"", 0},
	{"+", 0},
	{"20000000000000000000", 2147483647},
	{"  0000000000012345678", 12345678},
	{"000000000000000000", 0},
}

func Test(t *testing.T) {
	for i, tt := range tests {
		t.Run(strconv.Itoa(i+1), func(t *testing.T) {
			res := myAtoi(tt.in)
			if !reflect.DeepEqual(res, tt.out) {
				t.Errorf("with [%v] | got [%v] | want [%v]", tt.in, res, tt.out)
			}
		})
	}
}
