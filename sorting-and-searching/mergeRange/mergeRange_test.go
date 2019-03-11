package main

import (
	"reflect"
	"strconv"
	"testing"
)

var tests = []struct {
	in  []Interval
	out []Interval
}{
	//{[]Interval{{1,4},{0,0}},[]Interval{{0,0},{1,4}}},
	{[]Interval{{1, 4}, {4, 5}}, []Interval{{1, 5}}},
}

func Test(t *testing.T) {
	for i, tt := range tests {
		t.Run(strconv.Itoa(i+1), func(t *testing.T) {
			res := merge(tt.in)
			if !reflect.DeepEqual(res, tt.out) {
				t.Errorf("with [%v] | got [%v] | want [%v]", tt.in, res, tt.out)
			}
		})
	}
}
