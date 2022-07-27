package main

// 232. 用栈实现队列
// https://leetcode.cn/problems/implement-queue-using-stacks/

type MyQueue struct {
	in []int
	ou []int
}

func Constructor() MyQueue {
	return MyQueue{
		in: make([]int, 0),
		ou: make([]int, 0),
	}
}

func (this *MyQueue) Push(x int) {
	this.in = append(this.in, x)
}

func (this *MyQueue) Pop() int {
	this.adjust()
	if len(this.ou) == 0 {
		return -1
	}
	x := this.ou[len(this.ou)-1]
	this.ou = this.ou[:len(this.ou)-1]
	return x
}

func (this *MyQueue) adjust() {
	if len(this.ou) == 0 {
		for i := len(this.in) - 1; i >= 0; i-- {
			this.ou = append(this.ou, this.in[i])
		}
		this.in = make([]int, 0)
	}
}

func (this *MyQueue) Peek() int {
	this.adjust()
	if len(this.ou) == 0 {
		return -1
	}
	return this.ou[len(this.ou)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.in) == 0 && len(this.ou) == 0
}
