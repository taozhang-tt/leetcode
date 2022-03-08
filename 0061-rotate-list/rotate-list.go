package rotatelist

// 61. 旋转链表

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}
	root, tail, cnt := head, head, 1
	for tail.Next != nil { // 统计节点个数，同时tail移动到最后一个节点
		cnt, tail = cnt+1, tail.Next
	}
	k %= cnt // 循环
	if k > 0 {
		k = cnt - k // 向右移动k个位置，等于把前(cnt-k)个节点拼接到链表末尾
		for k > 1 {
			head, k = head.Next, k-1
		}
		tail.Next, root, head.Next = root, head.Next, nil
	}
	return root
}

func List2Slice(head *ListNode) []int {
	ret := make([]int, 0)
	for head != nil {
		ret = append(ret, head.Val)
		head = head.Next
	}
	return ret
}
