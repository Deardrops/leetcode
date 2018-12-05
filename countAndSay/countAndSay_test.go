package main

import (
	"reflect"
	"strconv"
	"testing"
)

var tests = []struct {
	in  int
	out string
}{
	{1, "1"},
	{2, "11"},
	{3, "21"},
	{4, "1211"},
}

func Test(t *testing.T) {
	for i, tt := range tests {
		t.Run(strconv.Itoa(i+1), func(t *testing.T) {
			res := countAndSay(tt.in)
			if !reflect.DeepEqual(res, tt.out) {
				t.Errorf("with [%v] | got [%v] | want [%v]", tt.in, res, tt.out)
			}
		})
	}
}
