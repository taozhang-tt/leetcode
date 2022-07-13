package main

// 142. 环形链表 II
// https://leetcode.cn/problems/linked-list-cycle-ii/

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	set := make(map[*ListNode]struct{})
	for head != nil {
		if _, ok := set[head]; ok {
			return head
		}
		set[head] = struct{}{}
		head = head.Next
	}
	return nil
}

func detectCycle2(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast, slow = fast.Next.Next, slow.Next
		if fast == slow {
			break
		}
	}
	if fast == nil || fast.Next == nil {
		return nil
	}
	fast = head
	for fast != slow {
		fast, slow = fast.Next, slow.Next
	}
	return fast
}
