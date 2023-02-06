package main

// [剑指 Offer 06. 从尾到头打印链表](https://leetcode.cn/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/)
// tag: 链表、翻转链表
// tip: 翻转链表、递归

type ListNode struct {
	Val  int
	Next *ListNode
}

// 暴力求解
func reversePrint(head *ListNode) []int {
	var (
		ret     = make([]int, 0)
		reverse func(head *ListNode)
	)
	reverse = func(head *ListNode) {
		if head == nil {
			return
		}
		reverse(head.Next)
		ret = append(ret, head.Val)
	}
	reverse(head)
	return ret
}

// 先翻转再遍历
func reverse(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		head.Next, pre, head = pre, head, head.Next
	}
	return pre
}

func reverseInRecursion(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	root := reverseInRecursion(head.Next)
	head.Next.Next = head
	head.Next = nil

	return root
}

func reversePrint2(head *ListNode) []int {
	head = reverseInRecursion(head)
	ret := make([]int, 0)
	for head != nil {
		ret = append(ret, head.Val)
		head = head.Next
	}
	return ret
}
