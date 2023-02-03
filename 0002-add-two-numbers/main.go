package main

// https://leetcode.cn/problems/add-two-numbers/
// 2. 两数相加
// 给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
// 请你将两个数相加，并以相同形式返回一个表示和的链表。
// 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	vir, sign := new(ListNode), 0
	head := vir

	for l1 != nil || l2 != nil {
		total := sign
		sign = 0

		if l1 != nil {
			total += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			total += l2.Val
			l2 = l2.Next
		}

		if total >= 10 {
			sign, total = 1, total-10
		}
		head.Next = &ListNode{
			Val: total,
		}
		head = head.Next
	}

	if sign > 0 {
		head.Next = &ListNode{
			Val: sign,
		}
	}
	return vir.Next
}
