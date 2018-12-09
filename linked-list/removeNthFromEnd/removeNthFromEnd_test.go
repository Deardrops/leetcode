package main

import (
	"reflect"
	"strconv"
	"testing"
)

type in struct {
	array []int
	n     int
}

var tests = []struct {
	in  in
	out []int
}{
	{in{[]int{1, 2, 3, 4, 5}, 2}, []int{1, 2, 3, 5}},
	{in{[]int{1}, 1}, []int{}},
	{in{[]int{1, 2}, 2}, []int{2}},
}

func Test(t *testing.T) {
	for i, tt := range tests {
		t.Run(strconv.Itoa(i+1), func(t *testing.T) {
			linkList := ArrayToLinkList(tt.in.array)
			linkList = removeNthFromEnd(linkList, tt.in.n)
			res := LinkListToArray(linkList)
			if !reflect.DeepEqual(res, tt.out) {
				t.Errorf("with [%v] | got [%v] | want [%v]", tt.in, res, tt.out)
			}
		})
	}
}

func ArrayToLinkList(array []int) *ListNode {
	var head *ListNode
	var curr *ListNode
	for _, num := range array {
		if head == nil {
			head = &ListNode{num, nil}
			curr = head
			continue
		}
		curr.Next = &ListNode{num, nil}
		curr = curr.Next
	}
	return head
}

func LinkListToArray(head *ListNode) []int {
	if head == nil {
		return make([]int, 0)
	}
	array := []int{head.Val}
	curr := head.Next
	for ; curr != nil; curr = curr.Next {
		array = append(array, curr.Val)
	}
	return array
}
