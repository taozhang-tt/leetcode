package main

// 92. 反转链表 II
// https://leetcode.cn/problems/reverse-linked-list-ii/
// 给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right
// 请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dumy := &ListNode{Next: head}
	prev := dumy
	for i := 1; i < left; i++ {
		prev = prev.Next
	}
	prev.Next = reverseK(prev.Next, right-left+1)
	return dumy.Next
}

// 反转链表
func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		head.Next, head, prev = prev, head.Next, head
	}
	return prev
}

// 反转链表的前K个节点, K 小于等于节点个数
func reverseK(head *ListNode, k int) *ListNode {
	var prev *ListNode
	root := head
	for k > 0 {
		head.Next, head, prev, k = prev, head.Next, head, k-1
	}
	root.Next = head
	return prev
}
