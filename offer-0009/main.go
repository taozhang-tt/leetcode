package main

// [剑指 Offer 09. 用两个栈实现队列](https://leetcode.cn/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof/)
// tag: 栈、队列

type MyStack struct {
	values []int
}

func StackConstructor() MyStack {
	return MyStack{make([]int, 0)}
}

func (this *MyStack) Push(x int) {
	this.values = append(this.values, x)
}

func (this *MyStack) Pop() int {
	if !this.IsEmpty() {
		x := this.values[len(this.values)-1]
		this.values = this.values[:len(this.values)-1]
		return x
	}

	return -1
}

func (this *MyStack) Top() int {
	if !this.IsEmpty() {
		return this.values[len(this.values)-1]
	}

	return -1
}

func (this *MyStack) IsEmpty() bool {
	return len(this.values) == 0
}

type CQueue struct {
	readStack  MyStack
	writeStack MyStack
}

func Constructor() CQueue {
	return CQueue{
		readStack:  StackConstructor(),
		writeStack: StackConstructor(),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.writeStack.Push(value)
}

func (this *CQueue) DeleteHead() int {
	if !this.readStack.IsEmpty() {
		return this.readStack.Pop()
	}

	for !this.writeStack.IsEmpty() {
		this.readStack.Push(this.writeStack.Pop())
	}

	return this.readStack.Pop()
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */

// type MyStack struct {
// 	values []int
// }

// func StackConstructor() MyStack {
// 	return MyStack{
// 		[]int{},
// 	}
// }

// func (this *MyStack) Push(value int) {
// 	this.values = append(this.values, value)
// }

// func (this *MyStack) Pop() int {
// 	if this.IsEmpty() {
// 		return -1
// 	}
// 	length := len(this.values)
// 	value := this.values[length-1]
// 	this.values = this.values[0 : length-1]
// 	return value
// }

// func (this *MyStack) IsEmpty() bool {
// 	return len(this.values) == 0
// }

// type CQueue struct {
// 	readStack  MyStack
// 	writeStack MyStack
// }

// func Constructor() CQueue {
// 	return CQueue{
// 		StackConstructor(),
// 		StackConstructor(),
// 	}
// }

// func (this *CQueue) AppendTail(value int) {
// 	this.writeStack.Push(value)
// }

// func (this *CQueue) DeleteHead() int {
// 	if !this.readStack.IsEmpty() {
// 		return this.readStack.Pop()
// 	}
// 	for !this.writeStack.IsEmpty() {
// 		this.readStack.Push(this.writeStack.Pop())
// 	}
// 	return this.readStack.Pop()
// }

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
