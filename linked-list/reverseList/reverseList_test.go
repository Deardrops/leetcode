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
	{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
	{[]int{}, []int{}},
	{[]int{1}, []int{1}},
}

func Test0(t *testing.T) {
	baseTest(reverseList, t)
}

func Test1(t *testing.T) {
	baseTest(recursiveReverseList, t)
}

func baseTest(f func(*ListNode) *ListNode, t *testing.T) {
	for i, tt := range tests {
		t.Run(strconv.Itoa(i+1), func(t *testing.T) {
			linkList := ArrayToLinkList(tt.in)
			linkList = f(linkList)
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
