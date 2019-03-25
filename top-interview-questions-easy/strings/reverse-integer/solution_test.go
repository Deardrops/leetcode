package main

import (
	"reflect"
	"strconv"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		in  int
		out int
	}{
		{123, 321},
		{-123, -321},
		{-2147483648, 0},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i+1), func(t *testing.T) {
			res := reverse(tt.in)
			if !reflect.DeepEqual(res, tt.out) {
				t.Errorf("with %q, got %q, want %q", tt.in, res, tt.out)
			}
		})
	}
}
