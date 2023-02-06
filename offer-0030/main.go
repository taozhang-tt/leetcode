package main

import "math"

// [剑指 Offer 30. 包含min函数的栈](https://leetcode.cn/problems/bao-han-minhan-shu-de-zhan-lcof/)
// tag: 栈
// tip: 辅助栈

type MinStack struct {
	values []int
	min    []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		make([]int, 0),
		[]int{math.MaxInt},
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func (this *MinStack) Push(x int) {
	this.values = append(this.values, x)
	this.min = append(this.min, min(this.Min(), x))
}

func (this *MinStack) Pop() {
	this.values = this.values[:len(this.values)-1]
	this.min = this.min[:len(this.min)-1]
}

func (this *MinStack) Top() int {
	return this.values[len(this.values)-1]
}

func (this *MinStack) Min() int {
	return this.min[len(this.min)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
