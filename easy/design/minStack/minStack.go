package main

import "fmt"

type MinStack struct {
	stack    []int
	minStack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	var stack, minStack []int
	return MinStack{stack, minStack}
}

func (this *MinStack) Push(x int) {
	length := len(this.minStack)
	if length == 0 || x <= this.minStack[length-1] {
		this.minStack = append(this.minStack, x)
	}
	this.stack = append(this.stack, x)
}

func (this *MinStack) Pop() {
	length := len(this.stack)
	length1 := len(this.minStack)
	if this.stack[length-1] == this.minStack[length1-1] {
		this.minStack = this.minStack[0 : length1-1]
	}
	this.stack = this.stack[0 : length-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
func main() {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	fmt.Println(obj.GetMin())
	obj.Pop()
	fmt.Println(obj.Top())
	fmt.Println(obj.GetMin())
}
