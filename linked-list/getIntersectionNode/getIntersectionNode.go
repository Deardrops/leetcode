package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 先将用两个指针（pA和pB）在两条链表上同时向后移动，当其中一个走完时，
// 将另外一个指针和它起点上的新指针（比如pB和newB）同时向后移动，直到pB走完
// 此时，pA指向A链表头，pB指向B链中间的某个位置，两者到其链表末尾的位置是一样的
// 将两者同时向后移动，相等的时候，该节点即为两条链表相交的节点
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	pA, newA, pB, newB := headA, headA, headB, headB

	for {
		if pA == nil {
			for pB != nil {
				pB = pB.Next
				newB = newB.Next
			}
			break
		}
		if pB == nil {
			for pA != nil {
				pA = pA.Next
				newA = newA.Next
			}
			break
		}
		pA = pA.Next
		pB = pB.Next
	}

	if newA == nil || newB == nil {
		return nil
	}
	if newA == newB {
		return newA
	}
	for newA.Next != nil && newB.Next != nil {
		newA = newA.Next
		newB = newB.Next
		if newA == newB {
			return newA
		}
	}
	return nil
}
